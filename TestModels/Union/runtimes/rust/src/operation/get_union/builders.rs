// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
pub use crate::operation::get_union::_get_union_output::GetUnionOutputBuilder;

pub use crate::operation::get_union::_get_union_input::GetUnionInputBuilder;

impl GetUnionInputBuilder {
    /// Sends a request with this input using the given client.
    pub async fn send_with(
        self,
        client: &crate::Client,
    ) -> ::std::result::Result<
        crate::operation::get_union::GetUnionOutput,
        crate::operation::get_union::GetUnionError,
    > {
        let mut fluent_builder = client.get_union();
        fluent_builder.inner = self;
        fluent_builder.send().await
    }
}
/// Fluent builder constructing a request to `GetUnion`.
///
#[derive(::std::clone::Clone, ::std::fmt::Debug)]
pub struct GetUnionFluentBuilder {
    client: crate::Client,
    inner: crate::operation::get_union::builders::GetUnionInputBuilder,
}
impl GetUnionFluentBuilder {
    /// Creates a new `GetUnionFluentBuilder`.
    pub(crate) fn new(client: crate::client::Client) -> Self {
        Self {
            client,
            inner: ::std::default::Default::default(),
        }
    }
    /// Access the GetUnion as a reference.
    pub fn as_input(&self) -> &crate::operation::get_union::builders::GetUnionInputBuilder {
        &self.inner
    }
    /// Sends the request and returns the response.
    pub async fn send(
        self,
    ) -> ::std::result::Result<
        crate::operation::get_union::GetUnionOutput,
        crate::operation::get_union::GetUnionError,
    > {
        let input = self
            .inner
            .build()
            // Using unhandled since GetUnion doesn't declare any validation,
            // and smithy-rs seems to not generate a ValidationError case unless there is
            // (but isn't that a backwards compatibility problem for output structures?)
            // Vanilla smithy-rs uses SdkError::construction_failure,
            // but we aren't using SdkError.
            .map_err(crate::operation::get_union::GetUnionError::unhandled)?;
        crate::operation::get_union::GetUnion::send(&self.client, input).await
    }

    #[allow(missing_docs)] // documentation missing in model
    pub fn value(mut self, input: crate::types::_my_union::MyUnion) -> Self {
        self.inner = self.inner.value(input);
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn set_value(
        mut self,
        input: ::std::option::Option<crate::types::_my_union::MyUnion>,
    ) -> Self {
        self.inner = self.inner.set_value(input);
        self
    }
    #[allow(missing_docs)] // documentation missing in model
    pub fn get_value(&self) -> &::std::option::Option<crate::types::_my_union::MyUnion> {
        self.inner.get_value()
    }
}
