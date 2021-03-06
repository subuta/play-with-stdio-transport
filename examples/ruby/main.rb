puts 'Try ctrl+c for exit.'

io = IO.popen("./server", "r+")

io.puts "{\"Type\":\"test\", \"Body\": {\"hoge\": \"fuga\"}}"

while line = io.gets
  puts line
end

# This will close process.
# io.close_write

# Wait for server process
# results tty opens forever because server process isn't a "daemon"
Process.wait(io.pid)