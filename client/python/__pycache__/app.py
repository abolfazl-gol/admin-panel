import grpc
from proto_pb2_grpc import ApiStub
from proto_pb2 import RegisterRequest

channel = grpc.insecure_channel('localhost:50051')
stub = ApiStub(channel)

user = stub.Register(RegisterRequest(email="aaaaaa", password="asdfasdf"))

print(user)
