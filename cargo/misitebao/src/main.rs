fn main() {
    let path = "https://misitebao.com";

    match open::that(path) {
        Ok(()) => println!("Opened '{}' successfully.", path),
        Err(err) => panic!("An error occurred when opening '{}': {}", path, err),
    }
}
