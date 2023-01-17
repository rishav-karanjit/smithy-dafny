namespace simple.types.integer

@aws.polymorph#localService(
  sdkId: "SimpleInteger",
  config: SimpleIntegerConfig,
)
service SimpleTypesInteger {
  version: "2021-11-01",
  resources: [],
  operations: [ GetInteger ],
  errors: [],
}

structure SimpleIntegerConfig {}

operation GetInteger {
  input: GetIntegerInput,
  output: GetIntegerOutput,
}

structure GetIntegerInput {
  value: Integer
}

structure GetIntegerOutput {
  value: Integer
}