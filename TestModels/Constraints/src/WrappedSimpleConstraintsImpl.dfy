// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
include "../Model/SimpleConstraintsTypesWrapped.dfy"

module {:extern "simpleconstraintsinternaldafnywrapped"} WrappedSimpleConstraintsService refines WrappedAbstractSimpleConstraintsService {
    import WrappedService = SimpleConstraints
    function method WrappedDefaultSimpleConstraintsConfig(): SimpleConstraintsConfig {
        SimpleConstraintsConfig
    }
}
