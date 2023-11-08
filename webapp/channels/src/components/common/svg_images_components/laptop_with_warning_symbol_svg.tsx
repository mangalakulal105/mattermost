// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from "react";

type SvgProps = {
    width?: number;
    height?: number;
};

const Svg = (props: SvgProps) => (
    <svg
        width={props.width?.toString() || "170"}
        height={props.height?.toString() || "129"}
        viewBox="0 0 170 129"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
    >
        <path
            d="M23.6258 81.7032H134.554C135.371 81.6915 136.149 81.3515 136.719 80.7574C137.29 80.1633 137.605 79.3636 137.597 78.5334V3.15128C137.605 2.32105 137.29 1.52137 136.719 0.927285C136.149 0.333195 135.371 -0.00690192 134.554 -0.0185547H23.6258C22.81 -0.00574534 22.0321 0.334704 21.4621 0.928539C20.8921 1.52237 20.5761 2.32139 20.5831 3.15128V78.5378C20.5773 79.3669 20.8938 80.1648 21.4637 80.7577C22.0336 81.3505 22.8107 81.6904 23.6258 81.7032Z"
            fill="#3F4350"
        />
        <path
            d="M11.3887 92.5439C11.3887 96.2965 14.4453 100.049 18.1524 100.049H140.031C143.525 100.049 146.791 96.3052 146.791 92.5439H11.3887Z"
            fill="#767D93"
        />
        <path
            d="M135.508 81.7031H22.6718L11.3887 92.5438H146.791L135.508 81.7031Z"
            fill="#D1D4DB"
        />
        <path
            d="M132.133 82.5371H26.0413L22.2539 87.1235H135.925L132.133 82.5371Z"
            fill="#AFB3C0"
        />
        <path
            d="M90.7217 88.7915H67.4624L65.7168 91.2932H92.4629L90.7217 88.7915Z"
            fill="#24262E"
        />
        <rect
            width="103.641"
            height="66.7116"
            transform="translate(27.2695 7.06982)"
            fill="white"
        />
        <rect
            x="27.2695"
            y="7.06982"
            width="104"
            height="67"
            fill="#3F4350"
            fillOpacity="0.16"
        />
        <path
            d="M79.0896 2.06641C79.4202 2.06641 79.7434 2.16423 80.0183 2.34748C80.2932 2.53074 80.5075 2.79121 80.634 3.09596C80.7605 3.40071 80.7936 3.73605 80.7291 4.05957C80.6646 4.38309 80.5054 4.68026 80.2717 4.9135C80.0379 5.14675 79.74 5.30559 79.4157 5.36994C79.0915 5.4343 78.7553 5.40125 78.4499 5.27502C78.1444 5.14879 77.8834 4.93504 77.6997 4.66077C77.516 4.38651 77.418 4.06405 77.418 3.7342C77.418 3.29187 77.5941 2.86766 77.9076 2.55489C78.2211 2.24212 78.6463 2.06641 79.0896 2.06641Z"
            fill="#989DAE"
        />
        <path
            d="M88.7203 97.5471H69.0285C68.3489 97.5471 66.5527 97.5471 66.5527 95.0454H91.2093C91.2093 97.5471 89.3602 97.5471 88.7203 97.5471Z"
            fill="#3F4350"
        />
        <path
            d="M60.1793 63.0923C57.6677 63.0923 56.5297 61.286 57.6504 59.0783L77.9502 19.2945C79.0997 17.0926 80.9101 17.0926 82.0366 19.2945L102.331 59.0783C103.48 61.2803 102.331 63.0923 99.8019 63.0923H60.1793Z"
            fill="#FFBC1F"
        />
        <path
            d="M76.8061 34.0084L78.8924 47.9941C78.9125 48.2716 79.0372 48.5312 79.2414 48.7207C79.4456 48.9103 79.7141 49.0156 79.993 49.0156C80.2719 49.0156 80.5404 48.9103 80.7446 48.7207C80.9488 48.5312 81.0736 48.2716 81.0937 47.9941L83.18 34.0084C83.5593 28.5552 76.421 28.5552 76.8061 34.0084Z"
            fill="#2D3039"
        />
        <path
            d="M79.9903 50.523C80.6221 50.5241 81.2393 50.712 81.7641 51.063C82.2888 51.4141 82.6976 51.9124 82.9385 52.4951C83.1795 53.0777 83.242 53.7186 83.118 54.3367C82.9941 54.9547 82.6893 55.5223 82.2421 55.9676C81.795 56.4129 81.2256 56.716 80.6059 56.8385C79.9862 56.9611 79.344 56.8976 78.7604 56.6561C78.1768 56.4147 77.6781 56.006 77.3272 55.4818C76.9763 54.9577 76.7891 54.3415 76.7891 53.7112C76.7891 53.292 76.8718 52.8769 77.0328 52.4898C77.1937 52.1026 77.4297 51.7509 77.727 51.4547C78.0244 51.1586 78.3773 50.9239 78.7657 50.764C79.1541 50.6041 79.5702 50.5222 79.9903 50.523Z"
            fill="#2D3039"
        />
    </svg>
);

export default Svg;
