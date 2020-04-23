import subprocess
from subprocess import Popen, PIPE, STDOUT

command = ["./client/client", "/home/bhegde/go/src/GoChat/client/pub_key", "/home/bhegde/go/src/GoChat/client/priv_key"]

process = subprocess.Popen(command, stdout=PIPE, stdin=PIPE, stderr=STDOUT)

print process.stdout.read()

print process.communicate("~~")[0]

print process.stdout.readline()
print process.stdout.readline()
print process.stdout.readline()