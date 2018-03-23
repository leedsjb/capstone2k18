import React from "react";

import Label from "../components/Label";
import TextInput from "../components/TextInput";

export default function renderField({
    input,
    label,
    type,
    meta: { touched, error, warning }
}) {
    return (
        <div>
            <Label>{label}</Label>
            <div>
                <TextInput {...input} placeholder={label} type={type} />
                {touched &&
                    ((error && <span>{error}</span>) ||
                        (warning && <span>{warning}</span>))}
            </div>
        </div>
    );
}
