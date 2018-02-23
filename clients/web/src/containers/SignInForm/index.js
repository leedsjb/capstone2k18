import React from "react";
import { Field, reduxForm } from "redux-form";

import renderField from "../../utils/renderField";
import validate from "./validate";
import warn from "./warn";

const SignInForm = props => {
    const { handleSubmit, onSubmit, submitting } = props;

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <Field
                name="email"
                type="email"
                component={renderField}
                label="Email"
            />
            <Field
                name="password"
                type="password"
                component={renderField}
                label="Password"
            />
            <button type="submit" disabled={submitting}>
                Submit
            </button>
        </form>
    );
};

export default reduxForm({
    form: "signInForm",
    validate,
    warn
})(SignInForm);
