import grpc
import hello_pb2
import hello_pb2_grpc

with grpc.insecure_channel('localhost:10000') as channel:
    stub = hello_pb2_grpc.HelloStub(channel)
    stub.Log(hello_pb2.String(msg="Hello, world!"))
