use super::*;

#[test]
fn test_empty_attrs() {
    let attrs = NixAttrs::construct(0, vec![]).expect("empty attr construction should succeed");

    assert!(
        matches!(attrs, NixAttrs::Empty),
        "empty attribute set should use optimised representation"
    );
}

#[test]
fn test_simple_attrs() {
    let attrs = NixAttrs::construct(
        1,
        vec![Value::String("key".into()), Value::String("value".into())],
    )
    .expect("simple attr construction should succeed");

    assert!(
        matches!(attrs, NixAttrs::Map(_)),
        "simple attribute set should use map representation",
    )
}

#[test]
fn test_kv_attrs() {
    let name_val = Value::String("name".into());
    let value_val = Value::String("value".into());
    let meaning_val = Value::String("meaning".into());
    let forty_two_val = Value::Integer(42);

    let kv_attrs = NixAttrs::construct(
        2,
        vec![
            value_val.clone(),
            forty_two_val.clone(),
            name_val.clone(),
            meaning_val.clone(),
        ],
    )
    .expect("constructing K/V pair attrs should succeed");

    match kv_attrs {
        NixAttrs::KV { name, value } if name == meaning_val || value == forty_two_val => {}

        _ => panic!(
            "K/V attribute set should use optimised representation, but got {:?}",
            kv_attrs
        ),
    }
}
