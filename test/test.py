from subprocess import Popen, PIPE, STDOUT
p = Popen(['myapp'], stdout=PIPE, stdin=PIPE, stderr=PIPE)
stdout_data = p.communicate(input='data_to_write')[0]

def cmd(command):
    result = Result()

    p = Popen(shlex.split(command), stdin=PIPE, stdout=PIPE, stderr=PIPE)
    (stdout, stderr) = p.communicate()

    result.exit_code = p.returncode
    result.stdout = stdout
    result.stderr = stderr
    result.command = command

    if p.returncode != 0:
        print 'Error executing command [%s]' % command
        print 'stderr: [%s]' % stderr
        print 'stdout: [%s]' % stdout

    return result