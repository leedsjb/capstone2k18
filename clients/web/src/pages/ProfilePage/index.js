import React from "react";
import { Helmet } from "react-helmet";

import TitleBar from "../../components/TitleBar";
import NavBar from "../../components/NavBar";
import Container from "../../components/Container";
import TabBar from "../../components/TabBar";
import FlexFillVH from "../../components/FlexFillVH";
import ScrollView from "../../components/ScrollView";
import ProfileAvatar from "../../components/ProfileAvatar";
import ButtonIcon from "../../components/ButtonIcon";

import ProfileProvider from "../../containers/ProfileProvider";

const ProfilePage = () => {
    return (
        <FlexFillVH flexDirection="column">
            <Helmet>
                <title>Profile</title>
            </Helmet>

            <TitleBar title="People" />
            <NavBar />

            <ScrollView>
                <Container>
                    <ProfileProvider
                        render={({ profile: { pending, data } }) => {
                            if (pending || Array.isArray(data)) {
                                return <div>Loading...</div>;
                            }

                            return (
                                <div>
                                    <ProfileAvatar
                                        fName={data.fName}
                                        size={96}
                                    />
                                    <div>{`${data.fName} ${data.lName}`}</div>
                                    <div>{data.position}</div>
                                    <ButtonIcon />
                                    <ButtonIcon />
                                    <ButtonIcon />
                                </div>
                            );
                        }}
                    />
                </Container>
            </ScrollView>

            <TabBar />
        </FlexFillVH>
    );
};

export default ProfilePage;
