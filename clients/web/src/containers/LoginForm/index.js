import React from "react";
import { Field, reduxForm } from "redux-form";

const LoginForm = props => {
    return (
        <form>
            <p>Login form</p>
        </form>
    );
};

export default reduxForm({
    form: "loginForm"
})(LoginForm);
