// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(missing_docs)] // documentation missing in model
#[non_exhaustive]
#[derive(::std::clone::Clone, ::std::cmp::PartialEq, ::std::fmt::Debug)]
pub struct GetIntegerKnownValueInput {
    #[allow(missing_docs)] // documentation missing in model
    pub value: ::std::option::Option<i32>,
}
impl GetIntegerKnownValueInput {
    #[allow(missing_docs)] // documentation missing in model
    pub fn value(&self) -> ::std::option::Option<i32> {
        self.value
    }
}
impl GetIntegerKnownValueInput {
    /// Creates a new builder-style object to manufacture [`GetIntegerKnownValueInput`](crate::operation::operation::GetIntegerKnownValueInput).
    pub fn builder(
    ) -> crate::operation::get_integer_known_value::builders::GetIntegerKnownValueInputBuilder {
        crate::operation::get_integer_known_value::builders::GetIntegerKnownValueInputBuilder::default()
    }
}

/// A builder for [`GetIntegerKnownValueInput`](crate::operation::operation::GetIntegerKnownValueInput).
#[non_exhaustive]
#[derive(
    ::std::clone::Clone, ::std::cmp::PartialEq, ::std::default::Default, ::std::fmt::Debug,
)]
pub struct GetIntegerKnownValueInputBuilder {
    pub(crate) value: ::std::option::Option<i32>,
}
impl GetIntegerKnownValueInputBuilder {
    #[allow(missing_docs)] // documentation missing in model
    pub fn value(mut self, input: impl ::std::convert::Into<i32>) -> Self {
        self.value = ::std::option::Option::Some(input.into());
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_value(mut self, input: ::std::option::Option<i32>) -> Self {
        self.value = input;
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_value(&self) -> &::std::option::Option<i32> {
        &self.value
    }
    /// Consumes the builder and constructs a [`GetIntegerKnownValueInput`](crate::operation::operation::GetIntegerKnownValueInput).
    pub fn build(
        self,
    ) -> ::std::result::Result<
        crate::operation::get_integer_known_value::GetIntegerKnownValueInput,
        ::aws_smithy_types::error::operation::BuildError,
    > {
        ::std::result::Result::Ok(
            crate::operation::get_integer_known_value::GetIntegerKnownValueInput {
                value: self.value,
            },
        )
    }
}
