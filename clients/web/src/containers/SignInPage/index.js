import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { connect } from "react-redux";

import SignInForm from "../SignInForm";

import { signIn } from "../../actions/auth/actions";

class SignInPage extends Component {
    render() {
        const { signIn } = this.props;

        return (
            <div>
                <Helmet>
                    <title>Sign in</title>
                </Helmet>

                <SignInForm onSubmit={signIn} />
            </div>
        );
    }
}

const mapDispatchToProps = {
    signIn
};

export default connect(null, mapDispatchToProps)(SignInPage);
