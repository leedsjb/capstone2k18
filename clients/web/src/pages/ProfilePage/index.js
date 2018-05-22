import React, { Component } from "react";
import { connect } from "react-redux";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";

import TitleBar from "../../components/TitleBar";
import Label from "../../components/Label";
import NavBar from "../../components/NavBar";
import TabBar from "../../components/TabBar";
import Box from "../../components/Box";
import Checkbox from "../../components/Checkbox";
import FlexFillVH from "../../components/FlexFillVH";
import ScrollView from "../../components/ScrollView";
import ColoredAvatar from "../../components/ColoredAvatar";
import ButtonIcon from "../../components/ButtonIcon";
import Heading from "../../components/Heading";
import Span from "../../components/Span";
import Switch from "../../components/Switch";
import TextInput from "../../components/TextInput";
import Container from "../../components/Container";

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
            console.log(currUser.data);
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
                                <Label>UW NetID</Label>
                                <TextInput
                                    value={currUser.data.uwNetID}
                                    disabled
                                />
                            </Box>
                            <Box mt={3}>
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
                            <Heading is="h2" mt={6}>
                                Notification preferences
                            </Heading>
                            <Flex alignItems="baseline" mt={4}>
                                <Checkbox defaultChecked />{" "}
                                <Span>
                                    Text me when I get assigned to a mission
                                </Span>
                            </Flex>
                        </Box>
                    </Flex>
                </Container>
            );
        }
        return <div>Loading...</div>;
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
                    <FlexFillVH>
                        An error has occurred:{" "}
                        {this.props.profile.error.toString()}
                    </FlexFillVH>
                ) : (
                    <ScrollView>
                        <Container>{this.renderProfile()}</Container>
                    </ScrollView>
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
