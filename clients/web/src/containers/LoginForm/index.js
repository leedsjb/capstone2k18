import React from "react";
import { Field, reduxForm } from "redux-form";
import { Button } from "rebass";

import renderField from "../../utils/renderField";
import validate from "./validate";
import warn from "./warn";

const LoginForm = props => {
    const { handleSubmit, pristine, reset, submitting } = props;

    return (
        <form onSubmit={handleSubmit}>
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
            <Button type="submit" disabled={submitting}>
                Submit
            </Button>
        </form>
    );
};

export default reduxForm({
    form: "loginForm",
    validate,
    warn
})(LoginForm);
