import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { connect } from "react-redux";

import Heading from "../../components/Heading";
import Container from "../../components/Container";

import SignInForm from "../../containers/SignInForm";

import { signIn } from "../../actions/auth/actions";

class SignInPage extends Component {
    render() {
        const { signIn } = this.props;

        return (
            <div>
                <Helmet>
                    <title>Sign in</title>
                </Helmet>

                <Container>
                    <Heading>AirliftNW Elevate</Heading>
                    <SignInForm onSubmit={signIn} />
                </Container>
            </div>
        );
    }
}

const mapDispatchToProps = {
    signIn
};

export default connect(null, mapDispatchToProps)(SignInPage);
