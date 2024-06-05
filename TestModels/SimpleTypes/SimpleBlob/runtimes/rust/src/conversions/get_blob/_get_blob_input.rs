// Code generated by software.amazon.smithy.rust.codegen.smithy-rs. DO NOT EDIT.
#[allow(dead_code)]
pub fn to_dafny(
    value: crate::operation::get_blob::GetBlobInput,
) -> ::std::rc::Rc<::simple_blob_dafny::r#_simple_dtypes_dblob_dinternaldafny_dtypes::GetBlobInput>
{
    let dafny_value = match value.value {
        Some(v) => ::simple_blob_dafny::_Wrappers_Compile::Option::Some {
            value: ::dafny_runtime::Sequence::from_array(&v),
        },
        None => ::simple_blob_dafny::_Wrappers_Compile::Option::None {},
    };
    ::std::rc::Rc::new(::simple_blob_dafny::r#_simple_dtypes_dblob_dinternaldafny_dtypes::GetBlobInput::GetBlobInput {
    value: ::std::rc::Rc::new(dafny_value)
  })
}

#[allow(dead_code)]
pub fn from_dafny(
    dafny_value: ::std::rc::Rc<
        ::simple_blob_dafny::r#_simple_dtypes_dblob_dinternaldafny_dtypes::GetBlobInput,
    >,
) -> crate::operation::get_blob::GetBlobInput {
    let value = if matches!(
        dafny_value.value().as_ref(),
        ::simple_blob_dafny::_Wrappers_Compile::Option::Some { .. }
    ) {
        Some(
            ::std::rc::Rc::try_unwrap(dafny_value.value().Extract().to_array())
                .unwrap_or_else(|rc| (*rc).clone()),
        )
    } else if matches!(
        dafny_value.value().as_ref(),
        ::simple_blob_dafny::_Wrappers_Compile::Option::None { .. }
    ) {
        None
    } else {
        panic!("Unreachable")
    };

    crate::operation::get_blob::GetBlobInput { value }
}