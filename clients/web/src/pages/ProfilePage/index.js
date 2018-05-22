import React, { Component } from "react";
import { connect } from "react-redux";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";

import TitleBar from "../../components/TitleBar";
import NavBar from "../../components/NavBar";
import TabBar from "../../components/TabBar";
import FlexFillVH from "../../components/FlexFillVH";
import ScrollView from "../../components/ScrollView";
import ProfileAvatar from "../../components/ProfileAvatar";
import ButtonIcon from "../../components/ButtonIcon";
import Heading from "../../components/Heading";
import Switch from "../../components/Switch";
import TextInput from "../../components/TextInput";
import Container from "../../components/Container";

import { fetchProfile } from "../../actions/profile/actions";

class ProfilePage extends Component {
    componentDidMount() {
        this.props.fetchProfile();
    }

    renderProfile(currUser) {
        if (
            !this.props.profile.pending &&
            !Array.isArray(this.props.profile.data)
        ) {
            return (
                <Container>
                    <ProfileAvatar fName={currUser.data.fName} size={96} />

                    <Heading is="h1">{`${currUser.data.fName} ${
                        currUser.data.lName
                    }`}</Heading>
                    <div>{currUser.data.position}</div>

                    <Heading>Notification preferences</Heading>
                    <Switch checked={false} color="primary" />
                    <TextInput value="Test" disabled />
                    <TextInput value="Test" disabled />
                    <TextInput value="Test" disabled />
                    <TextInput value="Test" disabled />
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
                        <Container>
                            {this.renderProfile(this.props.profile)}
                        </Container>
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
