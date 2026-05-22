import socket
import threading

def handle_client(client_socket, client_address):
    try:
        # Получаем запрос от клиента
        request = client_socket.recv(4096)
        
        if not request:
            client_socket.close()
            return

        # Разбираем запрос для получения хоста
        first_line = request.split(b'\n')[0]
        method, url, version = first_line.split(b' ')
        
        # Извлекаем хост из URL
        if url.startswith(b'http://'):
            url = url[7:]  # убираем http://
        
        # Разделяем хост и путь
        if b'/' in url:
            host = url.split(b'/')[0]
            path = b'/' + b'/'.join(url.split(b'/')[1:])
        else:
            host = url
            path = b'/'
        
        # Отделяем порт если есть
        if b':' in host:
            target_host, target_port = host.split(b':')
            target_port = int(target_port)
        else:
            target_host = host
            target_port = 80
        
        # Создаем соединение с целевым сервером
        server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server_socket.connect((target_host.decode(), target_port))
        
        # Пересылаем запрос на целевой сервер
        new_request = request.replace(url, path, 1)
        server_socket.send(new_request)
        
        # Получаем и пересылаем ответ клиенту
        while True:
            response = server_socket.recv(4096)
            if not response:
                break
            client_socket.send(response)
        
        server_socket.close()
        
    except Exception as e:
        print(f"Ошибка: {e}")
    finally:
        client_socket.close()

def start_proxy(host='127.0.0.1', port=8888):
    proxy_server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    proxy_server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    proxy_server.bind((host, port))
    proxy_server.listen(100)
    
    print(f"Прокси сервер запущен на {host}:{port}")
    
    try:
        while True:
            client_socket, client_address = proxy_server.accept()
            print(f"Подключен клиент: {client_address}")
            
            # Запускаем обработку клиента в отдельном потоке
            client_thread = threading.Thread(
                target=handle_client, 
                args=(client_socket, client_address)
            )
            client_thread.daemon = True
            client_thread.start()
            
    except KeyboardInterrupt:
        print("\n Process is stoped")
    finally:
        proxy_server.close()

if __name__ == "__main__":
    start_proxy()
