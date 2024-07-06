use std::path::PathBuf;

use clap::Parser;
use tracing::Level;

#[derive(Parser, Clone)]
pub struct Args {
    /// A global log level to use when printing logs.
    /// It's also possible to set `RUST_LOG` according to
    /// `tracing_subscriber::filter::EnvFilter`, which will always have
    /// priority.
    #[arg(long, default_value_t=Level::INFO)]
    pub log_level: Level,

    /// Path to a script to evaluate
    pub script: Option<PathBuf>,

    #[clap(long, short = 'E')]
    pub expr: Option<String>,

    /// Dump the raw AST to stdout before interpreting
    #[clap(long, env = "TVIX_DISPLAY_AST")]
    pub display_ast: bool,

    /// Dump the bytecode to stdout before evaluating
    #[clap(long, env = "TVIX_DUMP_BYTECODE")]
    pub dump_bytecode: bool,

    /// Trace the runtime of the VM
    #[clap(long, env = "TVIX_TRACE_RUNTIME")]
    pub trace_runtime: bool,

    /// Capture the time (relative to the start time of evaluation) of all events traced with
    /// `--trace-runtime`
    #[clap(long, env = "TVIX_TRACE_RUNTIME_TIMING", requires("trace_runtime"))]
    pub trace_runtime_timing: bool,

    /// Only compile, but do not execute code. This will make Tvix act
    /// sort of like a linter.
    #[clap(long)]
    pub compile_only: bool,

    /// Don't print warnings.
    #[clap(long)]
    pub no_warnings: bool,

    /// A colon-separated list of directories to use to resolve `<...>`-style paths
    #[clap(long, short = 'I', env = "NIX_PATH")]
    pub nix_search_path: Option<String>,

    /// Print "raw" (unquoted) output.
    #[clap(long)]
    pub raw: bool,

    /// Strictly evaluate values, traversing them and forcing e.g.
    /// elements of lists and attribute sets before printing the
    /// return value.
    #[clap(long)]
    pub strict: bool,

    #[arg(long, env, default_value = "memory://")]
    pub blob_service_addr: String,

    #[arg(long, env, default_value = "memory://")]
    pub directory_service_addr: String,

    #[arg(long, env, default_value = "memory://")]
    pub path_info_service_addr: String,

    #[arg(long, env, default_value = "dummy://")]
    pub build_service_addr: String,
}
