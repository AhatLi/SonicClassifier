// assets
import { IconDashboard, IconDeviceAnalytics } from '@tabler/icons';

// constant
const icons = {
    IconDashboard: IconDashboard,
    IconDeviceAnalytics
};

//-----------------------|| DASHBOARD MENU ITEMS ||-----------------------//

export const main = {
    id: 'main',
    title: 'main',
    type: 'group',
    children: [
        {
            id: 'default',
            title: 'main',
            type: 'item',
            url: '/main/default',
            icon: icons['IconDashboard'],
            breadcrumbs: false
        }
    ]
};
