from subprocess import Popen, PIPE, STDOUT

'''
p = Popen(['python', 'prog.py'], stdout=PIPE, stdin=PIPE, stderr=STDOUT)
#print p.stdout.readline().rstrip()
print p.communicate('mike 123')[0].rstrip()
'''

#command = ['./client/client', '/home/bhegde/go/src/GoChat/client/pub_key', '/home/bhegde/go/src/GoChat/client/priv_key']
command = ['go', 'run', 'userinput.go']
p = Popen(command, stdout=PIPE, stdin=PIPE, stderr=STDOUT)

print p.stdout.readline().rstrip()
print p.communicate("Hello")[0].rstrip()

print p.stdout.readline().rstrip()
print p.communicate("Hello")[0].rstrip()

#print p.communicate("World")[0].rstrip()

#print p.stdout.readline().rstrip()

#print p.communicate("Hello")[0].rstrip()