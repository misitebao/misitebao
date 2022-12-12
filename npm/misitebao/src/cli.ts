#! /usr/bin/env node
import open from "open";

const path: string = "https://misitebao.com";

(function main() {
  open(path)
    .then((res) => {
      console.log(`Opened '${path}' successfully.`);
    })
    .catch((err) => {
      console.log(`An error occurred when opening '${path}': ${err}`);
    });
})();
