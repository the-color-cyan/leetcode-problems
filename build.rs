use std::{env, fs, path::PathBuf};

fn main() {
    println!("cargo:rerun-if-changed=code");

    let mut problems: Vec<_> = fs::read_dir("code")
        .expect("read code directory")
        .filter_map(Result::ok)
        .map(|entry| entry.path())
        .filter(|path| path.extension().is_some_and(|ext| ext == "rs"))
        .collect();
    problems.sort();

    let modules = problems
        .iter()
        .enumerate()
        .map(|(index, path)| {
            let path = path.canonicalize().expect("resolve problem path");
            format!(
                "#[allow(dead_code)]\nmod problem_{index} {{\n    pub struct Solution;\n    include!({path:?});\n}}\n"
            )
        })
        .collect::<String>();

    fs::write(
        PathBuf::from(env::var_os("OUT_DIR").expect("OUT_DIR")).join("problems.rs"),
        modules,
    )
    .expect("write generated problem modules");
}
