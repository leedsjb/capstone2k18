import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { connect } from "react-redux";

import Heading from "../../components/Heading";
import Text from "../../components/Text";
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
                    <Text>Sign in with your AMC account</Text>
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
