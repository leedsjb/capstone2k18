import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Container, Heading } from "rebass";

import SignInForm from "../SignInForm";

import StyledMeasure from "../../components/StyledMeasure";

class SignInPage extends Component {
    render() {
        return (
            <div>
                <Helmet>
                    <title>Sign in</title>
                </Helmet>
                <Container>
                    <StyledMeasure mx="auto">
                        <Heading is="h1">Sign in</Heading>
                        <SignInForm
                            handleSubmit={values => {
                                console.log(values);
                            }}
                        />
                    </StyledMeasure>
                </Container>
            </div>
        );
    }
}

export default SignInPage;
