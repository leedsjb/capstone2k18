import React, { Component } from "react";

import Tab from "../../components/Tab";

const Tabs = props => {
    return (
        <div>
            {React.Children.map(props.children, (child, i) => {
                return (
                    <Tab
                        active={child.key === props.active ? true : false}
                        onClick={() => {
                            props.onChange(child.key);
                        }}
                    >
                        {child}
                    </Tab>
                );
            })}
        </div>
    );
};

export default Tabs;
