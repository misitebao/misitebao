#! /usr/bin/env node

const open = require("open");

const path = "https://misitebao.com";

(function main() {
  open(path)
    .then((res) => {
      console.log(`Opened '${path}' successfully.`);
    })
    .catch((err) => {
      console.log(`An error occurred when opening '${path}': ${err}`);
    });
})();
