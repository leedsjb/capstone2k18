import React, { Component } from "react";
import { connect } from "react-redux";
import { Helmet } from "react-helmet";

import TitleBar from "../../components/TitleBar";
import NavBar from "../../components/NavBar";
import Container from "../../components/Container";
import TabBar from "../../components/TabBar";
import FlexFillVH from "../../components/FlexFillVH";
import ScrollView from "../../components/ScrollView";
import ProfileAvatar from "../../components/ProfileAvatar";
import ButtonIcon from "../../components/ButtonIcon";

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
                <div>
                    <ProfileAvatar fName={currUser.data.fName} size={96} />
                    <div>{`${currUser.data.fName} ${currUser.data.lName}`}</div>
                    <div>{currUser.data.position}</div>
                    <ButtonIcon />
                    <ButtonIcon />
                    <ButtonIcon />
                </div>
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
