use bytes::Bytes;
use data_encoding::BASE64;
use thiserror::Error;

#[derive(PartialEq, Eq, Hash, Debug)]
pub struct B3Digest(Bytes);

// TODO: allow converting these errors to crate::Error
#[derive(Error, Debug)]
pub enum Error {
    #[error("invalid digest length: {0}")]
    InvalidDigestLen(usize),
}

impl B3Digest {
    // returns a copy of the inner [Vec<u8>].
    pub fn to_vec(&self) -> Vec<u8> {
        self.0.to_vec()
    }
}

impl From<B3Digest> for bytes::Bytes {
    fn from(val: B3Digest) -> Self {
        val.0
    }
}

impl TryFrom<Vec<u8>> for B3Digest {
    type Error = Error;

    // constructs a [B3Digest] from a [Vec<u8>].
    // Returns an error if the digest has the wrong length.
    fn try_from(value: Vec<u8>) -> Result<Self, Self::Error> {
        if value.len() != 32 {
            Err(Error::InvalidDigestLen(value.len()))
        } else {
            Ok(Self(value.into()))
        }
    }
}

impl TryFrom<bytes::Bytes> for B3Digest {
    type Error = Error;

    // constructs a [B3Digest] from a [bytes::Bytes].
    // Returns an error if the digest has the wrong length.
    fn try_from(value: bytes::Bytes) -> Result<Self, Self::Error> {
        if value.len() != 32 {
            Err(Error::InvalidDigestLen(value.len()))
        } else {
            Ok(Self(value))
        }
    }
}

impl From<&[u8; 32]> for B3Digest {
    fn from(value: &[u8; 32]) -> Self {
        Self(value.to_vec().into())
    }
}

impl Clone for B3Digest {
    fn clone(&self) -> Self {
        Self(self.0.to_owned())
    }
}

impl std::fmt::Display for B3Digest {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "b3:{}", BASE64.encode(&self.0))
    }
}
