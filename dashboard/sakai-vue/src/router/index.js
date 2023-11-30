import { createRouter, createWebHashHistory } from 'vue-router';
import AppLayout_Customer from '@/layout/AppLayout_Customer.vue';
import AppLayout_Driver from '@/layout/AppLayout_Driver.vue';
import store from '@/store';
const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: '/customer',
            // With sidebar & topbar
            component: AppLayout_Customer,
            children: [
                {
                    path: '/',
                    name: 'dashboard',
                    component: () => import('@/views/Dashboard.vue')
                },
                {
                    path: '/logo',
                    name: 'logo',
                    component: () => import('@/views/UI/logo.vue')
                },
                {
                    path: '/register',
                    name: 'register',
                    component: () => import('@/views/UI/register.vue')
                },
                {
                    path: '/login',
                    name: 'UIlogin',
                    component: () => import('@/views/UI/login.vue')
                },
                {
                    path: '/profile',
                    name: 'profile',
                    component: () => import('@/views/UI/profile.vue')
                },
                {
                    path: '/uikit/formlayout',
                    name: 'formlayout',
                    component: () => import('@/views/uikit/FormLayout.vue')
                },
                {
                    path: '/uikit/input',
                    name: 'input',
                    component: () => import('@/views/uikit/Input.vue')
                },
                {
                    path: '/uikit/floatlabel',
                    name: 'floatlabel',
                    component: () => import('@/views/uikit/FloatLabel.vue')
                },
                {
                    path: '/uikit/invalidstate',
                    name: 'invalidstate',
                    component: () => import('@/views/uikit/InvalidState.vue')
                },
                {
                    path: '/uikit/button',
                    name: 'button',
                    component: () => import('@/views/uikit/Button.vue')
                },
                {
                    path: '/uikit/table',
                    name: 'table',
                    component: () => import('@/views/uikit/Table.vue')
                },
                {
                    path: '/uikit/list',
                    name: 'list',
                    component: () => import('@/views/uikit/List.vue')
                },
                {
                    path: '/uikit/tree',
                    name: 'tree',
                    component: () => import('@/views/uikit/Tree.vue')
                },
                {
                    path: '/uikit/panel',
                    name: 'panel',
                    component: () => import('@/views/uikit/Panels.vue')
                },

                {
                    path: '/uikit/overlay',
                    name: 'overlay',
                    component: () => import('@/views/uikit/Overlay.vue')
                },
                {
                    path: '/uikit/media',
                    name: 'media',
                    component: () => import('@/views/uikit/Media.vue')
                },
                {
                    path: '/uikit/menu',
                    component: () => import('@/views/uikit/Menu.vue'),
                    children: [
                        {
                            path: '/uikit/menu',
                            component: () => import('@/views/uikit/menu/PersonalDemo.vue')
                        },
                        {
                            path: '/uikit/menu/seat',
                            component: () => import('@/views/uikit/menu/SeatDemo.vue')
                        },
                        {
                            path: '/uikit/menu/payment',
                            component: () => import('@/views/uikit/menu/PaymentDemo.vue')
                        },
                        {
                            path: '/uikit/menu/confirmation',
                            component: () => import('@/views/uikit/menu/ConfirmationDemo.vue')
                        }
                    ]
                },
                {
                    path: '/uikit/message',
                    name: 'message',
                    component: () => import('@/views/uikit/Messages.vue')
                },
                {
                    path: '/uikit/file',
                    name: 'file',
                    component: () => import('@/views/uikit/File.vue')
                },
                {
                    path: '/uikit/charts',
                    name: 'charts',
                    component: () => import('@/views/uikit/Chart.vue')
                },
                {
                    path: '/uikit/misc',
                    name: 'misc',
                    component: () => import('@/views/uikit/Misc.vue')
                },
                {
                    path: '/blocks',
                    name: 'blocks',
                    component: () => import('@/views/utilities/Blocks.vue')
                },
                {
                    path: '/utilities/icons',
                    name: 'icons',
                    component: () => import('@/views/utilities/Icons.vue')
                },
                {
                    path: '/pages/timeline',
                    name: 'timeline',
                    component: () => import('@/views/pages/Timeline.vue')
                },
                {
                    path: '/pages/empty',
                    name: 'empty',
                                    },
                {                   path: '/customer/home',
                    name: 'CustomerHome',
                    component: () => import('@/views/pages/Empty.vue')
                },
                {
                    path: '/customer/history',
                    name: 'CustomerHistory',
                    component: () => import('@/views/pages/Empty.vue')
                },
                {
                    path: '/customer/search', // /search?q=yourSearchQuery
                    name: 'CustomerSearch',
                    component: () => import('@/views/pages/Empty.vue')
                },
                {
                    path: '/customer/setting',
                    name: 'CutomerSetting',
                    component: () => import('@/views/pages/Empty.vue')
                },
                {
                    path: '/customer/mytrip',
                    name: 'MyTrip',
                    component: () => import('@/views/pages/customer/MyTrip.vue')
                },
                {
                    path: '/customer/favorite',
                    name: 'Favorite',
                    component: () => import('@/views/pages/customer/Favorite.vue')
                }
            ]
        },
        {
            path: '/driver',
            // With sidebar & topbar
            component: AppLayout_Driver,
            children: [
                {
                    path: '/driver/home',
                    name: 'DriverHome',
                    component: () => import('@/views/pages/Empty.vue')
                },
                {
                    path: '/driver/history',
                    name: 'DriverHistory',
                    component: () => import('@/views/pages/Empty.vue')
                },
                {
                    path: '/driver/create',
                    name: 'DriverCreate',
                    component: () => import('@/views/pages/Empty.vue')
                },
                {
                    path: '/driver/setting',
                    name: 'DriverSetting',
                    component: () => import('@/views/pages/Empty.vue')
                },
                {
                    path: '/driver/profile',
                    name: 'DriverProfile',
                    component: () => import('@/views/pages/driver/DriverProfile.vue')
                },
                {
                    path: '/driver/applylist',
                    name: 'ApplyList',
                    component: () => import('@/views/pages/driver/ApplyList.vue')
                },
            ]
        },
        {
            path: '/auth/login',
            name: 'login',
            component: () => import('@/views/pages/auth/Login.vue')
        },
        {
            path: '/auth/access',
            name: 'accessDenied',
            component: () => import('@/views/pages/auth/Access.vue')
        },
        {
            path: '/auth/register',
            name: 'Register',
            component: () => import('@/views/pages/auth/Access.vue')
        },
        {
            path: '/chat/:userId',
            name: 'Chatroom',
            component: () => import('@/views/pages/utils/Chat.vue')
        },
        {
            path: '/onTrip/:tripId',
            name: 'OnTrip',
            component: () => import('@/views/pages/utils/OnTrip.vue')
        },
        {
            path: '/TripDetail/:tripId',
            name: 'TripDetail',
            component: () => import('@/views/pages/utils/TripDetail.vue')
        }
    ],
});


// Navigation guard to check authentication before each navigation
router.beforeEach(async (to, from, next) => {
    if (to.name !== 'login' && !store.getters.login) {
        try {
            console.log('check login');
            await store.dispatch("checkLogin");
            next();
        } catch {
            next({ name: "login" });
        }
    } else {
    console.log('have login');
      next(); // Continue with the navigation
    }
});

export default router;
