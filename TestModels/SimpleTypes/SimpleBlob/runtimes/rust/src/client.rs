// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.

use aws_smithy_types::error::operation::BuildError;

#[derive(::std::clone::Clone, ::std::fmt::Debug)]
pub struct Client {
    pub(crate) dafny_client: ::dafny_runtime::Object<dyn ::simple_blob_dafny::r#_simple_dtypes_dblob_dinternaldafny_dtypes::ISimpleTypesBlobClient>
}

impl Client {
    /// Creates a new client from the service [`Config`](crate::Config).
    #[track_caller]
    pub fn from_conf(
        conf: crate::types::simple_blob_config::SimpleBlobConfig,
    ) -> Result<Self, BuildError> {
        let inner = ::simple_blob_dafny::_simple_dtypes_dblob_dinternaldafny::_default::SimpleBlob(
            &crate::conversions::simple_blob_config::_simple_blob_config::to_dafny(conf),
        );
        if matches!(
            inner.as_ref(),
            ::simple_blob_dafny::_Wrappers_Compile::Result::Failure { .. }
        ) {
            // TODO: convert error - the potential types are not modeled!
            return Err(BuildError::other(
                ::aws_smithy_types::error::metadata::ErrorMetadata::builder()
                    .message("Invalid client config")
                    .build(),
            ));
        }
        Ok(Self {
            dafny_client: ::dafny_runtime::UpcastTo::<dafny_runtime::Object<(dyn ::simple_blob_dafny::r#_simple_dtypes_dblob_dinternaldafny_dtypes::ISimpleTypesBlobClient + 'static)>>::upcast_to(inner.Extract()),
        })
    }
}

mod get_blob;

mod get_blob_known_value;
