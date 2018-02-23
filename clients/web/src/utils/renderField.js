import React from "react";

export default function renderField({
    input,
    label,
    type,
    meta: { touched, error, warning }
}) {
    return (
        <div>
            <label>{label}</label>
            <div>
                <input {...input} placeholder={label} type={type} />
                {touched &&
                    ((error && <span>{error}</span>) ||
                        (warning && <span>{warning}</span>))}
            </div>
        </div>
    );
}
