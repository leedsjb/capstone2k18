import React from "react";
import { Field, reduxForm } from "redux-form";

import renderField from "../../utils/renderField";
import validate from "./validate";
import warn from "./warn";

import Button from "../../components/Button";
import Box from "../../components/Box";

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
            <Box mt={4}>
                <Button type="submit" disabled={submitting}>
                    Sign in
                </Button>
            </Box>
        </form>
    );
};

export default reduxForm({
    form: "signInForm",
    validate,
    warn
})(SignInForm);
