import { spawn } from 'child_process'
import path from 'path'

const child = spawn(path.resolve(__dirname, './server'))

console.log('Try ctrl+c for exit.')

child.stdin.setEncoding('utf-8');
child.stdout.pipe(process.stdout);

child.stdin.write(`{"Type":"test", "Body": {"hoge": "fuga"}}\n`);