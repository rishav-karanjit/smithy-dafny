include "../Model/ExampleSimpletypesTypes.dfy"

module SimpleTypesImpl refines AbstractExampleSimpletypesOperations  {
  datatype Config = Config
  type InternalConfig = Config
  predicate ValidInternalConfig?(config: InternalConfig)
  {true}
  function ModifiesInternalConfig(config: InternalConfig) : set<object>
  {{}}
   predicate GetStringEnsuresPublicly(input: GetStringInput, output: Result<GetStringOutput, Error>) {
    true
   }
 method GetString ( config: InternalConfig,  input: GetStringInput )
 returns (output: Result<GetStringOutput, Error>) {
    var res := GetStringOutput(result := input.stringValue);
    return Success(res);
 }
}