import { createServer } from 'http';
import { readFile } from 'fs/promises';
import { join } from 'path';
import { createReadStream } from 'fs';

const PORT = 3000;

async function run() {
  const jsonPath = join('..', 'common', 'json', `${process.argv[2]}.json`);
  const jsonData = await readFile(jsonPath);
  const json = JSON.parse(jsonData);

  const filePath = join('..', 'common', 'files', `${process.argv[2]}.txt`);

  createServer((req, res) => {
    if (req.url === '/json-stringify') {
      res.end(JSON.stringify(json));
    }
    if (req.url === '/download-files') {
      createReadStream(filePath).pipe(res);
    }

  }).listen(PORT, () => { console.log(`up on ${PORT}`) })

}

run()