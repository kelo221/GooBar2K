type TabContent = {
    title: string;
    text: string;
};

type Tab = {
    id: string;
    label: string;
    content: TabContent;
};

type BottomStats = {
    codec: string;
    bitrate: string;
    frequency: string;
    channelType: string;
    playtime: string;
};

type ProgressBarState = {
    progress: number;
};
