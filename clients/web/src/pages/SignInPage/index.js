import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";
import { connect } from "react-redux";
import { Link } from "react-router-dom";

import Heading from "../../components/Heading";
import Text from "../../components/Text";
import Container from "../../components/Container";
import Box from "../../components/Box";
import Logo from "../../components/Logo";

import SignInForm from "../../containers/SignInForm";

import { signIn } from "../../actions/auth/actions";

class SignInPage extends Component {
    render() {
        const { signIn } = this.props;

        return (
            <Flex flexDirection="column" flex={1}>
                <Helmet>
                    <title>Sign in</title>
                </Helmet>

                <Box p={3}>
                    <Flex justifyContent="space-between">
                        <Link to="/">
                            <Logo />
                        </Link>
                        <Link to="/signin">Sign in</Link>
                    </Flex>
                </Box>

                <Flex flexDirection="column" flex={1} justifyContent="center">
                    <Container>
                        <Heading>AirliftNW Elevate</Heading>
                        <Text>Sign in with your AMC account</Text>
                        <SignInForm onSubmit={signIn} />
                    </Container>
                </Flex>
            </Flex>
        );
    }
}

const mapDispatchToProps = {
    signIn
};

export default connect(null, mapDispatchToProps)(SignInPage);