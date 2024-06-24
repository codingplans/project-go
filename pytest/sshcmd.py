import paramiko


def ssh_command(hostname, command):
    # 创建SSH对象
    ssh = paramiko.SSHClient()
    # 允许连接不在know_hosts文件中的主机
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())

    try:
        # 连接服务器
        # ssh.connect(hostname=hostname,username="zhangzhy",key_filename="/Users/darren/.ssh/id_rsa.zzy-m1-hj")
        ssh.connect(hostname=hostname,username="zhangzhy")
        # 执行命令
        stdin, stdout, stderr = ssh.exec_command(command)
        # 获取命令输出
        output = stdout.read().decode()
        # 打印输出结果
        print(output)
    except Exception as e:
        print(f"An error occurred: {e}")
    finally:
        # 关闭连接
        ssh.close()


# 执行命令
hostname = '192.168.0.205'
command = 'ls'
ssh_command(hostname, command)