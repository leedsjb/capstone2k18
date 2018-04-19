import React from "react";
import { Helmet } from "react-helmet";

import TitleBar from "../../components/TitleBar";
import NavBar from "../../components/NavBar";
import Container from "../../components/Container";
import TabBar from "../../components/TabBar";
import FlexFullHeight from "../../components/FlexFullHeight";
import ScrollView from "../../components/ScrollView";

import ProfileProvider from "../../containers/ProfileProvider";

const ProfilePage = () => {
    return (
        <FlexFullHeight flexDirection="column">
            <Helmet>
                <title>Profile</title>
            </Helmet>

            <TitleBar title="People" />
            <NavBar />

            <ScrollView>
                <Container>
                    <ProfileProvider
                        render={({ profile: { pending, data } }) => {
                            if (pending) {
                                return <div>Loading...</div>;
                            }

                            return (
                                <div>
                                    <div>{`${data.fName} ${data.lName}`}</div>
                                    <div>{data.position}</div>
                                </div>
                            );
                        }}
                    />
                </Container>
            </ScrollView>

            <TabBar />
        </FlexFullHeight>
    );
};

export default ProfilePage;
