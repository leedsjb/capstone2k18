import React, { Component } from "react";
import { connect } from "react-redux";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";

import Box from "../../components/Box";
import ColoredAvatar from "../../components/ColoredAvatar";
import Container from "../../components/Container";
import FlexFillVH from "../../components/FlexFillVH";
import Heading from "../../components/Heading";
import Label from "../../components/Label";
import NavBar from "../../components/NavBar";
import ScrollView from "../../components/ScrollView";
import Span from "../../components/Span";
import LoadingSpinner from "../../components/LoadingSpinner";
import TabBar from "../../components/TabBar";
import TextInput from "../../components/TextInput";
import TitleBar from "../../components/TitleBar";
import Error from "../../components/Error";

import { fetchProfile } from "../../actions/profile/actions";

class ProfilePage extends Component {
    componentDidMount() {
        this.props.fetchProfile();
    }

    renderProfile() {
        let currUser = this.props.profile;

        if (
            !this.props.profile.pending &&
            !Array.isArray(this.props.profile.data)
        ) {
            return (
                <Container py={12}>
                    <Flex alignItems="center" flexDirection="column">
                        <ColoredAvatar fName={currUser.data.fName} size={96} />
                        <Box maxWidth="30em" w={1}>
                            <Heading is="h2" mt={6}>
                                Account
                            </Heading>
                            <Box mt={2}>
                                <Span>
                                    Contact Chris Bishop if you'd like to update
                                    your information
                                </Span>
                            </Box>
                            <Box mt={8}>
                                <Label>First name</Label>
                                <TextInput
                                    value={currUser.data.fName}
                                    disabled
                                />
                            </Box>
                            <Box mt={3}>
                                <Label>Last name</Label>
                                <TextInput
                                    value={currUser.data.lName}
                                    disabled
                                />
                            </Box>
                            <Box mt={3}>
                                <Label>Email</Label>
                                <TextInput
                                    value={currUser.data.email}
                                    disabled
                                />
                            </Box>
                            <Box mt={3}>
                                <Label>Phone</Label>
                                <TextInput
                                    value={currUser.data.mobile}
                                    disabled
                                />
                            </Box>
                            <Box mt={3}>
                                <Label>Position</Label>
                                <TextInput
                                    value={currUser.data.position}
                                    disabled
                                />
                            </Box>
                            <Box mt={3}>
                                <Label>Special quals</Label>
                                <TextInput
                                    value={
                                        currUser.data.specialQuals
                                            ? currUser.data.specialQuals
                                            : "N/A"
                                    }
                                    disabled
                                />
                            </Box>
                        </Box>
                    </Flex>
                </Container>
            );
        }
        return <LoadingSpinner />;
    }

    render() {
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Profile</title>
                </Helmet>

                <TitleBar title="Profile" />
                <NavBar />
                {this.props.profile.error ? (
                    <Flex
                        flexDirection="column"
                        flex={1}
                        alignItems="center"
                        justifyContent="center"
                    >
                        <Error
                            title="An error has occurred"
                            content={this.props.profile.error.toString()}
                        />
                    </Flex>
                ) : (
                    <ScrollView>{this.renderProfile()}</ScrollView>
                )}
                <TabBar />
            </FlexFillVH>
        );
    }
}

function mapStateToProps(state, ownProps) {
    return {
        profile: state.profile
    };
}

const mapDispatchToProps = {
    fetchProfile
};

export default connect(mapStateToProps, mapDispatchToProps)(ProfilePage);
