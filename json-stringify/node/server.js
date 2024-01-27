import { createServer } from 'http';
import { readFile } from 'fs/promises';
import { join } from 'path';

const PORT = 3000;

async function run() {
  const path = join('..', 'common', `${process.argv[2]}.json`);
  const data = await readFile(path);
  const json = JSON.parse(data);

  createServer((req, res) => {
    res.end(JSON.stringify(json))
  }).listen(PORT, () => { console.log(`up on ${PORT}`) })

}

run()