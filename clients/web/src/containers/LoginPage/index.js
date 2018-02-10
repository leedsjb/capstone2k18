import React, { Component } from "react";
import { Heading } from "rebass";

import LoginForm from "../LoginForm";

import StyledMeasure from "../../components/StyledMeasure";

class LoginPage extends Component {
    render() {
        return (
            <StyledMeasure mx="auto">
                <Heading is="h1">Sign in</Heading>
                <LoginForm
                    handleSubmit={values => {
                        console.log(values);
                    }}
                />
            </StyledMeasure>
        );
    }
}

export default LoginPage;
