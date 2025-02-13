import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import path from "path";

const PROTO_PATH = path.join(
  __dirname,
  "protos",
  "test.proto"
);
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition) as any;
const testApiProto = protoDescriptor.test.api.proto;
const TestService = testApiProto.TestService;
const client = new TestService(
  "localhost:50051",
  grpc.credentials.createInsecure()
);
const requestData = {
  collection_1: [1, 3, 5, 7],
  collection_2: [2, 4, 6, 8],
  collection_3: [20, 15, 10, 0],
};
client.Merge(requestData, (err: grpc.ServiceError, response: any) => {
  if (err) {
    console.error("Error calling Merge:", err);
  } else {
    console.log("Merge response:", response);
  }
});
