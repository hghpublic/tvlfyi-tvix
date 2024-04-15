//! This module contains all the data structures used to track information
//! about inodes, which present tvix-castore nodes in a filesystem.
use std::time::Duration;

use bytes::Bytes;

use crate::proto as castorepb;
use crate::B3Digest;

#[derive(Clone, Debug)]
pub enum InodeData {
    Regular(B3Digest, u64, bool),  // digest, size, executable
    Symlink(bytes::Bytes),         // target
    Directory(DirectoryInodeData), // either [DirectoryInodeData:Sparse] or [DirectoryInodeData:Populated]
}

/// This encodes the two different states of [InodeData::Directory].
/// Either the data still is sparse (we only saw a [castorepb::DirectoryNode],
/// but didn't fetch the [castorepb::Directory] struct yet, or we processed a
/// lookup and did fetch the data.
#[derive(Clone, Debug)]
pub enum DirectoryInodeData {
    Sparse(B3Digest, u64),                                  // digest, size
    Populated(B3Digest, Vec<(u64, castorepb::node::Node)>), // [(child_inode, node)]
}

impl InodeData {
    /// Constructs a new InodeData by consuming a [Node].
    /// It splits off the orginal name, so it can be used later.
    pub fn from_node(node: castorepb::node::Node) -> (Self, Bytes) {
        match node {
            castorepb::node::Node::Directory(n) => (
                Self::Directory(DirectoryInodeData::Sparse(
                    n.digest.try_into().unwrap(),
                    n.size,
                )),
                n.name,
            ),
            castorepb::node::Node::File(n) => (
                Self::Regular(n.digest.try_into().unwrap(), n.size, n.executable),
                n.name,
            ),
            castorepb::node::Node::Symlink(n) => (Self::Symlink(n.target), n.name),
        }
    }

    pub fn as_fuse_file_attr(&self, inode: u64) -> fuse_backend_rs::abi::fuse_abi::Attr {
        fuse_backend_rs::abi::fuse_abi::Attr {
            ino: inode,
            // FUTUREWORK: play with this numbers, as it affects read sizes for client applications.
            blocks: 1024,
            size: match self {
                InodeData::Regular(_, size, _) => *size,
                InodeData::Symlink(target) => target.len() as u64,
                InodeData::Directory(DirectoryInodeData::Sparse(_, size)) => *size,
                InodeData::Directory(DirectoryInodeData::Populated(_, ref children)) => {
                    children.len() as u64
                }
            },
            mode: match self {
                InodeData::Regular(_, _, false) => libc::S_IFREG | 0o444, // no-executable files
                InodeData::Regular(_, _, true) => libc::S_IFREG | 0o555,  // executable files
                InodeData::Symlink(_) => libc::S_IFLNK | 0o444,
                InodeData::Directory(_) => libc::S_IFDIR | 0o555,
            },
            ..Default::default()
        }
    }

    pub fn as_fuse_entry(&self, inode: u64) -> fuse_backend_rs::api::filesystem::Entry {
        fuse_backend_rs::api::filesystem::Entry {
            inode,
            attr: self.as_fuse_file_attr(inode).into(),
            attr_timeout: Duration::MAX,
            entry_timeout: Duration::MAX,
            ..Default::default()
        }
    }

    /// Returns the u32 fuse type
    pub fn as_fuse_type(&self) -> u32 {
        #[allow(clippy::let_and_return)]
        let ty = match self {
            InodeData::Regular(_, _, _) => libc::S_IFREG,
            InodeData::Symlink(_) => libc::S_IFLNK,
            InodeData::Directory(_) => libc::S_IFDIR,
        };
        // libc::S_IFDIR is u32 on Linux and u16 on MacOS
        #[cfg(target_os = "macos")]
        let ty = ty as u32;

        ty
    }
}
