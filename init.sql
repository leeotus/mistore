SET NAMES utf8mb4;

SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for access
-- ----------------------------
DROP TABLE IF EXISTS `access`;

CREATE TABLE `access` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `module_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0',
    `type` tinyint(1) NULL DEFAULT NULL,
    `action_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `module_id` int(0) NULL DEFAULT NULL,
    `sort` int(0) NULL DEFAULT NULL,
    `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `add_time` int(0) NULL DEFAULT NULL,
    `status` int(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 114 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of access
-- ----------------------------
INSERT INTO
    `access`
VALUES (
        52,
        '管理员管理',
        1,
        '',
        '',
        0,
        105,
        '管理员管理',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        53,
        '角色管理',
        1,
        '',
        '',
        0,
        100,
        '角色管理',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        54,
        '管理员管理',
        2,
        '管理员列表',
        'manager',
        52,
        100,
        '管理员列表',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        55,
        '管理员管理',
        2,
        '增加管理员',
        'manager/add',
        52,
        101,
        '管理员列表',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        56,
        '管理员管理',
        3,
        '编辑管理员',
        'manager/edit',
        52,
        100,
        '编辑管理员',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        57,
        '管理员管理',
        3,
        '删除管理员',
        'manager/delete',
        52,
        100,
        '删除管理员',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        59,
        '角色管理',
        2,
        '角色列表',
        'role',
        53,
        100,
        '角色列表',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        60,
        '角色管理',
        2,
        '增加角色',
        'role/add',
        53,
        102,
        '增加角色',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        61,
        '角色管理',
        3,
        '编辑角色',
        'role/edit',
        53,
        100,
        '编辑角色',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        62,
        '角色管理',
        3,
        '删除角色',
        'role/delete',
        53,
        100,
        '删除角色',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        63,
        '权限管理',
        1,
        '',
        '',
        0,
        100,
        '权限管理',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        64,
        '权限管理',
        2,
        '权限列表',
        'access',
        63,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        67,
        '权限管理',
        2,
        '增加权限',
        'access/add',
        63,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        69,
        '系统设置',
        2,
        '轮播图列表',
        'focus',
        106,
        101,
        '1111',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        70,
        '系统设置',
        3,
        '增加轮播图',
        'focus/add',
        106,
        100,
        '增加轮播图',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        71,
        '系统设置',
        3,
        '编辑轮播图',
        'focus/edit',
        106,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        75,
        '系统设置',
        3,
        '删除轮播图',
        'focus/delete',
        106,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        76,
        '管理员管理',
        3,
        '执行增加管理员',
        'manager/doAdd',
        52,
        100,
        '执行增加',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        77,
        '管理员管理',
        3,
        '执行修改管理员',
        'manager/doEdit',
        52,
        100,
        '执行修改',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        78,
        '角色管理',
        3,
        '执行增加角色',
        'role/doAdd',
        53,
        100,
        '执行增加',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        79,
        '角色管理',
        3,
        '执行修改角色',
        'role/doEdit',
        53,
        100,
        '执行修改',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        80,
        '角色管理',
        3,
        '角色授权',
        'role/auth',
        53,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        81,
        '角色管理',
        3,
        '执行角色授权',
        'role/doAuth',
        53,
        100,
        '执行授权',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        82,
        '权限管理',
        3,
        '修改权限',
        'access/edit',
        63,
        100,
        '执行修改',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        83,
        '权限管理',
        3,
        '删除权限',
        'access/delete',
        63,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        84,
        '权限管理',
        3,
        '执行增加权限',
        'access/doAdd',
        63,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        85,
        '权限管理',
        3,
        '执行修改权限',
        'access/doEdit',
        63,
        100,
        '执行修改\r\n',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        86,
        '系统设置',
        3,
        '执行增加',
        'focus/doAdd',
        106,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        87,
        '商品管理',
        1,
        '',
        '',
        0,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        88,
        '商品管理',
        2,
        '商品分类列表',
        'goodsCate',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        89,
        '商品管理',
        3,
        '增加商品分类',
        'goodsCate/add',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        90,
        '商品管理',
        3,
        '执行增加商品分类',
        'goodsCate/doAdd',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        91,
        '商品管理',
        3,
        '修改商品分类',
        'goodsCate/edit',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        92,
        '商品管理',
        3,
        '执行修改商品分类',
        'goodsCate/doEdit',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        93,
        '商品管理',
        3,
        '删除商品分类',
        'goodsCate/delete',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        94,
        '商品管理',
        2,
        '商品类型列表',
        'goodsType',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        95,
        '商品管理',
        3,
        '增加商品类型',
        'goodsType/add',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        96,
        '商品管理',
        3,
        '编辑商品类型',
        'goodsType/edit',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        97,
        '商品管理',
        3,
        '执行增加 商品类型',
        'goodsType/doAdd',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        98,
        '商品管理',
        3,
        '执行修改 商品类型',
        'goodsType/doEdit',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        99,
        '商品管理',
        3,
        '删除 商品类型',
        'goodsType/delete',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        100,
        '商品管理',
        2,
        '商品列表',
        'goods',
        87,
        100,
        '商品列表',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        101,
        '商品管理',
        3,
        '增加商品',
        'goods/add',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        102,
        '商品管理',
        3,
        '执行 增加商品',
        'goods/doAdd',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        103,
        '商品管理',
        3,
        '修改商品',
        'goods/edit',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        104,
        '商品管理',
        3,
        '执行 修改商品',
        'goods/doEdit',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        105,
        '商品管理',
        3,
        '删除商品',
        'goods/delete',
        87,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        106,
        '系统设置',
        1,
        '',
        '',
        0,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        107,
        '系统设置',
        2,
        '导航管理',
        'nav',
        106,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        108,
        '系统设置',
        3,
        '增加导航',
        'nav/add',
        106,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        109,
        '系统设置',
        3,
        '编辑导航',
        'nav/edit',
        106,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        110,
        '系统设置',
        3,
        '删除导航',
        'nav/delete',
        106,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        111,
        '系统设置',
        3,
        '执行增加',
        'nav/doAdd',
        106,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        112,
        '系统设置',
        3,
        '执行修改',
        'nav/doEdit',
        106,
        100,
        '',
        0,
        1
    );

INSERT INTO
    `access`
VALUES (
        113,
        '系统设置',
        2,
        '商店设置',
        'setting',
        106,
        100,
        '',
        0,
        1
    );

-- ----------------------------
-- Table structure for focus
-- ----------------------------
DROP TABLE IF EXISTS `focus`;

CREATE TABLE `focus` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `focus_type` tinyint(1) NULL DEFAULT NULL,
    `focus_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `sort` int(0) NULL DEFAULT NULL,
    `status` tinyint(1) NULL DEFAULT NULL,
    `add_time` int(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of focus
-- ----------------------------
INSERT INTO
    `focus`
VALUES (
        14,
        '小米手机',
        1,
        'static/upload/20211101/1635757964474199700.jpg',
        'http://www.baidu.com',
        1115,
        1,
        1631677671
    );

INSERT INTO
    `focus`
VALUES (
        16,
        '小米电视1111',
        1,
        'static/upload/20211101/1635757979944161500.jpg',
        'http://a.baidu.com',
        1222,
        1,
        1631679244
    );

INSERT INTO
    `focus`
VALUES (
        17,
        'ces',
        1,
        'static/upload/20211101/1635758018523031700.jpg',
        'http://www.baidu.com',
        100,
        1,
        1635758011
    );

INSERT INTO
    `focus`
VALUES (
        18,
        '啊啊啊',
        1,
        'static/upload/20211102/1635817134986870600.png',
        'http://www.baidu.com',
        100,
        1,
        1635817134
    );

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;

CREATE TABLE `goods` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `sub_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `goods_sn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `cate_id` int(0) NULL DEFAULT NULL,
    `click_count` int(0) NULL DEFAULT NULL,
    `goods_number` int(0) NULL DEFAULT NULL,
    `price` decimal(10, 2) NULL DEFAULT NULL,
    `market_price` decimal(10, 2) NULL DEFAULT NULL,
    `relation_goods` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `goods_attr` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `goods_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `goods_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `goods_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `goods_gift` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `goods_fitting` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `goods_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `goods_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `goods_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
    `is_delete` tinyint(0) NULL DEFAULT NULL,
    `is_hot` tinyint(0) NULL DEFAULT NULL,
    `is_best` tinyint(0) NULL DEFAULT NULL,
    `is_new` tinyint(0) NULL DEFAULT NULL,
    `goods_type_id` int(0) NULL DEFAULT NULL,
    `sort` int(0) NULL DEFAULT NULL,
    `status` tinyint(0) NULL DEFAULT NULL,
    `add_time` int(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 54 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods
-- ----------------------------
INSERT INTO
    `goods`
VALUES (
        19,
        '小米9-8GB+256GB',
        '火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起',
        '',
        38,
        100,
        1000,
        999.00,
        999.00,
        '19,53',
        '尺寸:41,42,43',
        '1,2,4,5',
        '8GB+256GB',
        'static/upload/20200617/1592392307796676500.jpg',
        '1,2',
        '1,2',
        '',
        '',
        '<p>火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起</p><p><br></p><p><img src=\"http://bee.apiying.com/static/upload/20211101/1635739607166546900.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p><p><img src=\"http://bee.apiying.com/static/upload/20211101/1635740680831942900.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p><p><br></p>',
        0,
        0,
        1,
        1,
        1,
        100,
        1,
        1592392307
    );

INSERT INTO
    `goods`
VALUES (
        20,
        'Redmi Note 11 5G手机 1亿像素 55W有线闪充 50W无线闪充  6G+128GB 手机',
        '购机赠价值897元善诊3人体检套餐；赠Keep会员7日体验卡；光大信用卡分期支付满1000元减80元，数量有限',
        '',
        2,
        100,
        100,
        3699.00,
        4599.00,
        '20,54',
        '',
        '2,5,7,8',
        '6G+128GB',
        'static/upload/20211117/1637139107685884400.jpg',
        '',
        '',
        '',
        '',
        '<p>火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起</p>',
        0,
        1,
        1,
        0,
        1,
        0,
        1,
        1592392495
    );

INSERT INTO
    `goods`
VALUES (
        21,
        '小米8年度旗舰222',
        '火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起',
        '',
        36,
        100,
        1000,
        1112.00,
        1113.00,
        '1,2',
        '1,2',
        '2,3,4,5',
        '3GB+32GB',
        'static/upload/20211102/1635849810407008900.png',
        '1,2',
        '1,2',
        '1,2',
        '1,2',
        '<p>火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起</p><p><br></p><p><img src=\"http://bee.apiying.com/static/upload/20211101/1635736323217965200.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p>',
        0,
        0,
        1,
        1,
        1,
        11,
        1,
        1592392825
    );

INSERT INTO
    `goods`
VALUES (
        22,
        'Redmi 7A',
        '「3GB+32GB到手价仅549元」4000mAh超长续航 / 骁龙8核处理器 / 标配10W快充 / AI人脸解锁 / 大字体，大音量，无线收音机 / 整机生活防泼溅 / 极简模式，亲情守护',
        '',
        2,
        100,
        1000,
        549.00,
        799.00,
        '',
        '',
        '3,4',
        '3GB+32GB',
        'static/upload/20200622/1592820040.jpg',
        '',
        '',
        '',
        '',
        '<p><span style=\"color: rgb(51, 51, 51); font-family: F9ab65; font-size: 10.4922px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: left; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-style: initial; text-decoration-color: initial; display: inline !important; float: none;\">小巧机身蕴藏4000mAh大电量，配合MIUI系统级省电优化，精细调控，从此告别电量焦虑，尽情尽欢！</span></p>',
        0,
        0,
        1,
        0,
        1,
        100,
        1,
        1592820016
    );

INSERT INTO
    `goods`
VALUES (
        23,
        'Redmi 智能电视 X65',
        '全金属边框/4K超高清/MEMC运动补偿/8单元重低音音响系统',
        '',
        5,
        100,
        1000,
        2999.00,
        3299.00,
        '',
        '',
        '4',
        '56寸',
        'static/upload/20200622/1592820111.jpg',
        '',
        '',
        '',
        '',
        '<p><span style=\'color: rgb(176, 176, 176); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 14px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-style: initial; text-decoration-color: initial; display: inline !important; float: none;\'>全金属边框/4K超高清/MEMC运动补偿/8单元重低音音响系统</span></p>',
        0,
        0,
        1,
        0,
        0,
        100,
        0,
        1592820111
    );

INSERT INTO
    `goods`
VALUES (
        24,
        'RedmiBook 13 全面屏',
        '四窄边全面屏 / 全新十代酷睿™处理器 / 全金属超轻机身 / MX250 高性能独显 / 小米互传 / 专业「飓风」散热系统 / 11小时长续航',
        '',
        20,
        100,
        1000,
        4499.00,
        4799.00,
        '',
        '',
        '4,5',
        '8G+128G',
        'static/upload/20200622/1592820244.jpg',
        '',
        '',
        '',
        '',
        '<p><span style=\'color: rgb(176, 176, 176); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 14px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: start; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-style: initial; text-decoration-color: initial; display: inline !important; float: none;\'>四窄边全面屏 / 全新十代酷睿&trade;处理器 / 全金属超轻机身 / MX250 高性能独显 / 小米互传 / 专业「飓风」散热系统 / 11小时长续航</span> </p>',
        0,
        0,
        1,
        0,
        0,
        100,
        1,
        1592820244
    );

INSERT INTO
    `goods`
VALUES (
        25,
        '米家电磁炉',
        '99挡微调控火 / 支持低温烹饪 / 100+烹饪模式',
        '',
        1,
        100,
        1000,
        299.00,
        399.00,
        '',
        '',
        '',
        '',
        'static/upload/20200622/1592820331.jpg',
        '',
        '',
        '',
        '',
        '<p>米家电磁炉</p>',
        0,
        1,
        1,
        0,
        0,
        100,
        1,
        1592820331
    );

INSERT INTO
    `goods`
VALUES (
        26,
        '黑鲨双向快充移动电源',
        '18W双向快充 / 铠甲机身 / 一入三出 / 炫酷灯效',
        '',
        37,
        100,
        1000,
        0.00,
        0.00,
        '',
        '',
        '',
        '',
        'static/upload/20200622/1592820494.jpg',
        '',
        '',
        '',
        '',
        '',
        0,
        0,
        1,
        0,
        0,
        100,
        1,
        1592820494
    );

INSERT INTO
    `goods`
VALUES (
        36,
        '小米手机2222',
        '1111111111',
        '',
        1,
        100,
        0,
        4444.00,
        5555.00,
        '7777',
        '10',
        '1,3,5',
        '33332222222222',
        'static/upload/20211011/1633944450109113300.jpg',
        '8888',
        '999999999',
        '',
        '',
        '<p>小米10</p><p><img src=\"/static/upload/20211011/1633944295927657900.png\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p><p><img src=\"/static/upload/20211011/1633944470896453500.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p>',
        0,
        0,
        0,
        0,
        2,
        12,
        1,
        1633755416
    );

INSERT INTO
    `goods`
VALUES (
        37,
        '小米电视测试',
        '222222222222222222',
        '',
        1,
        100,
        0,
        444444.00,
        5555.00,
        '7777777777',
        '1000000011111100000',
        '1,2,3,4,5',
        '333333333333333',
        'static/upload/20211009/1633755741820253400.png',
        '8888888888',
        '999999',
        '111 1111111 111111111111111111',
        '121212',
        '<p>666666666666666</p>',
        0,
        1,
        1,
        1,
        2,
        0,
        0,
        1633755741
    );

INSERT INTO
    `goods`
VALUES (
        38,
        '小米手机测试111',
        '124214214214',
        '',
        1,
        100,
        0,
        0.00,
        0.00,
        '',
        '',
        '2,5',
        '',
        'static/upload/20211009/1633755959396859300.png',
        '',
        '',
        '',
        '',
        '',
        0,
        1,
        1,
        1,
        1,
        0,
        0,
        1633755959
    );

INSERT INTO
    `goods`
VALUES (
        39,
        'Redmi k30',
        '6.53\"水滴大屏 | 5020mAh超长续航 | G80高性能处理器 | 全场景 AI 四摄 | 大功率扬声器 | 指纹识别 | 人脸解锁 | 红外遥控',
        '',
        38,
        100,
        100,
        899.00,
        899.00,
        '',
        '',
        '',
        '',
        'static/upload/20211116/1637026344085801400.jpg',
        '',
        '',
        '',
        '',
        '',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1635502706
    );

INSERT INTO
    `goods`
VALUES (
        40,
        'Xiaomi MIX 4',
        'CUP全面屏 | 真彩原色 + 120Hz | 一体化轻量陶瓷机身 | 高通骁龙™888+ | WiFi 6 增强版 | 石墨烯「冰封」散热系统',
        '',
        37,
        100,
        100,
        0.00,
        0.00,
        '',
        '',
        '',
        '',
        'static/upload/20211116/1637026171480899500.jpg',
        '',
        '',
        '',
        '',
        '',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1635503000
    );

INSERT INTO
    `goods`
VALUES (
        41,
        'Xiaomi Civi',
        '轻薄潮流设计 | 丝绒AG工艺 | 原生美肌人像 | 像素级肌肤焕新技术 | 3200万高清质感自拍 | 双柔光灯+自动对焦 | 3D曲面OLED柔性屏 | 120Hz+Dolby Vision | 4500mAh 大电量 | 55W有线闪充 | 立体声双扬声器',
        '',
        36,
        100,
        100,
        1200.00,
        1400.00,
        '',
        '',
        '',
        '',
        'static/upload/20211116/1637026086634961500.jpg',
        '',
        '',
        '',
        '',
        '',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1635503077
    );

INSERT INTO
    `goods`
VALUES (
        42,
        'Redmi Note 10 5G',
        ' 5G小金刚｜旗舰长续航｜双5G待机｜5000mAh充电宝级大容量｜4800万高清相机｜天玑700八核高性能处理器',
        '',
        35,
        100,
        100,
        0.00,
        0.00,
        '',
        '',
        '',
        '',
        'static/upload/20211116/1637025991576339600.jpg',
        '',
        '',
        '',
        '',
        '',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1635503644
    );

INSERT INTO
    `goods`
VALUES (
        43,
        'Xiaomi 10S',
        '骁龙870 | 对称式双扬立体声 | 1亿像素 8K电影相机 | 33W有线快充 | 30W无线快充 | 10W反向充电 | 4780mAh超大电池 | LPDDR5+UFS3.0+Wi-Fi 6 | VC液冷散热 | 双模5G',
        '',
        35,
        100,
        100,
        2699.00,
        3699.00,
        '',
        '',
        '1,2,3',
        '8GB+128GB',
        'static/upload/20211102/1635841579767962200.jpg',
        '',
        '',
        '',
        '',
        '<p id=\"isPasted\"><br></p><p>高通骁龙&trade;870</p><p>哈曼卡顿｜对称式双扬立体声</p><p>4780mAh 大电量</p><p>三重快充 33W有线+30W无线+10W反向充电</p><p>小至尊经典外观</p><p>LPDDR5+UFS3.0+WiFi6</p><p>1 亿像素电影相机</p><p>8K 电影模式</p><p><br></p>',
        0,
        0,
        1,
        0,
        1,
        100,
        1,
        1635841578
    );

INSERT INTO
    `goods`
VALUES (
        44,
        'Xiaomi 11 Pro',
        '至高享24期免息，赠蓝牙耳机Air2 SE，+1元得30W立式无线充',
        '',
        2,
        100,
        100,
        0.00,
        0.00,
        '',
        '',
        '2,3,4',
        '',
        'static/upload/20211102/1635841908156579200.jpg',
        '',
        '',
        '',
        '',
        '<p><br></p><p id=\"isPasted\" style=\"text-align: center;\"><span style=\"font-size: 24px;\">联合研发18个月</span></p><p style=\"text-align: center;\"><span style=\"font-size: 24px;\">2亿影像投入，打造超强规格主摄</span></p><p style=\"text-align: center;\"><span style=\"font-size: 24px;\">这是颗&ldquo;巨型大底&rdquo;的面积，甚至可以媲美专业便携式相机，超大的进光量，</span></p><p style=\"text-align: center;\"><span style=\"font-size: 24px;\">带来了前所未有丰富的细节，&ldquo;夜视&rdquo;能力因此远超人眼，更能&ldquo;看懂&rdquo;夜色。</span></p><p><img src=\"/static/upload/20211102/1635841855622147000.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p><p><br></p>',
        0,
        0,
        1,
        0,
        0,
        100,
        1,
        1635841907
    );

INSERT INTO
    `goods`
VALUES (
        45,
        '小米移动电源3 20000mAh USB-C双向快充版',
        '',
        '',
        20,
        100,
        100,
        100.00,
        100.00,
        '',
        '',
        '',
        '',
        'static/upload/20211102/1635844763742258900.jpg',
        '',
        '',
        '',
        '',
        '',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1635844763
    );

INSERT INTO
    `goods`
VALUES (
        46,
        '小米移动电源3 10000mAh 超级闪充版 （50W）',
        '',
        '',
        20,
        100,
        100,
        125.00,
        155.00,
        '',
        '',
        '',
        '',
        'static/upload/20211102/1635844808324401400.jpg',
        '',
        '',
        '',
        '',
        '',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1635844808
    );

INSERT INTO
    `goods`
VALUES (
        47,
        '小米6A Type-C快充数据线',
        '',
        '',
        9,
        100,
        100,
        29.00,
        29.00,
        '',
        '',
        '',
        '',
        '',
        '',
        '',
        '',
        '',
        '',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1635845354
    );

INSERT INTO
    `goods`
VALUES (
        48,
        '小米USB-C数据线 编织线版 100cm',
        '',
        '',
        9,
        100,
        100,
        0.00,
        0.00,
        '',
        '',
        '',
        '',
        'static/upload/20211102/1635845426055325800.jpg',
        '',
        '',
        '',
        '',
        '<p><img src=\"/static/upload/20211102/1635845418913722200.png\" style=\"width: 300px;\" class=\"fr-fic fr-dib\"></p>',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1635845425
    );

INSERT INTO
    `goods`
VALUES (
        49,
        'Redmi Note 11 Pro系列',
        '三星AMOLED高刷屏 | JBL 对称式立体声 | 一亿像素超清影像 | 天玑920液冷芯 | VC液冷立体散热',
        '',
        2,
        100,
        100,
        0.00,
        0.00,
        '',
        '',
        '',
        '',
        'static/upload/20211116/1637025826328576500.jpg',
        '',
        '',
        '',
        '',
        '',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1637025826
    );

INSERT INTO
    `goods`
VALUES (
        50,
        '小米全面屏电视 55英寸PRO E55S',
        'Amlogic T972超强悍处理器 / 4K超高清画质 细腻如真 / 支持8K视频解码 / 2G+32G超大存储 / 内置小爱同学 语音控制更方便 / 智能Patchwall汇聚海量好内容',
        '',
        19,
        100,
        100,
        2399.00,
        2499.00,
        '',
        '',
        '',
        '',
        'static/upload/20211116/1637049463471284100.jpg',
        '',
        '',
        '',
        '',
        '',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1637049463
    );

INSERT INTO
    `goods`
VALUES (
        51,
        '米家互联网对开门冰箱 540L',
        '风冷无霜/环绕出风/纤薄箱体/电脑控温,持久保鲜/智能互联',
        '',
        13,
        100,
        100,
        2899.00,
        2999.00,
        '23,24,39',
        '',
        '1,2,3,4',
        '',
        'static/upload/20211116/1637049592911969300.jpg',
        '23,24,39',
        '23,24,39',
        '',
        '',
        '',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1637049592
    );

INSERT INTO
    `goods`
VALUES (
        52,
        '米家微波炉',
        '智能APP操控 / 平板式加热 / 专项分类解冻 / 20L容量 / 30+精选食谱',
        '',
        1,
        100,
        100,
        399.00,
        499.00,
        '',
        '',
        '',
        '',
        'static/upload/20211116/1637049679925704800.jpg',
        '',
        '',
        '',
        '',
        '',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1637049679
    );

INSERT INTO
    `goods`
VALUES (
        53,
        '小米9 6GB+128GB',
        '',
        '',
        38,
        100,
        100,
        1113.00,
        1167.00,
        '19,53',
        '',
        '2,6,7',
        '6GB+128GB',
        'static/upload/20211116/1637063708413624300.jpg',
        '',
        '',
        '',
        '',
        '<p id=\"isPasted\">火爆热卖中，6GB+64GB/6GB+128GB闪降100元，到手价1299元起</p><p><br></p><p><img src=\"http://bee.apiying.com/static/upload/20211101/1635739607166546900.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dii\"></p><p><img src=\"http://bee.apiying.com/static/upload/20211101/1635740680831942900.jpg\" style=\"width: 300px;\" class=\"fr-fic fr-dii\"></p>',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1637063708
    );

INSERT INTO
    `goods`
VALUES (
        54,
        'Redmi Note 11 5G手机 1亿像素 55W有线闪充 50W无线闪充  8G+256GB 手机',
        '双卡双5G | X轴线性马达 | 5000mAh 大电量 | 33W快充 | 立体声双扬声器 | 天玑810处理器 | 90Hz变速高刷屏',
        '',
        2,
        100,
        100,
        4199.00,
        4599.00,
        '20,54',
        '',
        '2,8',
        ' 8G+256GB',
        '',
        '',
        '',
        '',
        '',
        '<p id=\"isPasted\" style=\'margin: 0px; font-weight: bolder; font-size: 18px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'><br>8GB + 256GB 最高可选</p><p style=\'margin: 0px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 17px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'>4GB + 128GB</p><p style=\'margin: 0px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 17px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'>6GB + 128GB</p><p style=\'margin: 0px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 17px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'>8GB + 128GB</p><p style=\'margin: 0px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 17px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'>8GB + 256GB</p><p style=\'margin: 0px; color: rgb(0, 0, 0); font-family: \"Helvetica Neue\", Helvetica, Arial, \"Microsoft Yahei\", \"Hiragino Sans GB\", \"Heiti SC\", \"WenQuanYi Micro Hei\", sans-serif; font-size: 17px; font-style: normal; font-variant-ligatures: normal; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: justify; text-indent: 0px; text-transform: none; white-space: normal; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;\'>LPDDR4X 内存 +UFS2.2 闪存</p>',
        0,
        0,
        0,
        0,
        0,
        100,
        1,
        1637139500
    );

-- ----------------------------
-- Table structure for goods_attr
-- ----------------------------
DROP TABLE IF EXISTS `goods_attr`;

CREATE TABLE `goods_attr` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `goods_id` int(0) NULL DEFAULT NULL,
    `attribute_cate_id` int(0) NULL DEFAULT NULL,
    `attribute_id` int(0) NULL DEFAULT NULL,
    `attribute_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `attribute_type` tinyint(1) NULL DEFAULT NULL,
    `attribute_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `sort` int(0) NULL DEFAULT NULL,
    `add_time` int(0) NULL DEFAULT NULL,
    `status` tinyint(1) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 137 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_attr
-- ----------------------------
INSERT INTO
    `goods_attr`
VALUES (
        99,
        36,
        2,
        2,
        '主体',
        3,
        '1111',
        10,
        1634722805,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        100,
        36,
        2,
        3,
        '内存',
        1,
        'aaaaaaa',
        10,
        1634722805,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        101,
        36,
        2,
        4,
        '硬盘',
        1,
        'bbbb',
        10,
        1634722806,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        102,
        36,
        2,
        5,
        '显示器',
        1,
        'cccc',
        10,
        1634722806,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        103,
        36,
        2,
        6,
        '支持蓝牙',
        3,
        '是\r\n',
        10,
        1634722806,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        104,
        38,
        1,
        1,
        '基本信息',
        1,
        '124214',
        10,
        1634722815,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        105,
        38,
        1,
        7,
        '性能	',
        2,
        '214214',
        10,
        1634722815,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        106,
        38,
        1,
        8,
        '相机',
        2,
        '214214',
        10,
        1634722815,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        107,
        38,
        1,
        9,
        '支持蓝牙',
        3,
        '是\r\n',
        10,
        1634722815,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        108,
        37,
        2,
        2,
        '主体',
        3,
        '111\r\n',
        10,
        1634722866,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        109,
        37,
        2,
        3,
        '内存',
        1,
        '内存',
        10,
        1634722866,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        110,
        37,
        2,
        4,
        '硬盘',
        1,
        '硬盘',
        10,
        1634722866,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        111,
        37,
        2,
        5,
        '显示器',
        1,
        '显示器:',
        10,
        1634722866,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        112,
        37,
        2,
        6,
        '支持蓝牙',
        3,
        '否',
        10,
        1634722866,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        117,
        43,
        1,
        1,
        '基本信息',
        1,
        '高通骁龙™870',
        10,
        1637025929,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        118,
        43,
        1,
        7,
        '性能	',
        2,
        '4780mAh 大电量',
        10,
        1637025929,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        119,
        43,
        1,
        8,
        '相机',
        2,
        '4780mAh 大电量',
        10,
        1637025929,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        120,
        43,
        1,
        9,
        '支持蓝牙',
        3,
        '是\r\n',
        10,
        1637025929,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        133,
        19,
        1,
        1,
        '基本信息',
        1,
        '',
        10,
        1637064232,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        134,
        19,
        1,
        7,
        '性能	',
        2,
        '',
        10,
        1637064232,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        135,
        19,
        1,
        8,
        '相机',
        2,
        '',
        10,
        1637064232,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        136,
        19,
        1,
        9,
        '支持蓝牙',
        3,
        '是\r\n',
        10,
        1637064232,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        157,
        20,
        1,
        1,
        '基本信息',
        1,
        '',
        10,
        1637139540,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        158,
        20,
        1,
        7,
        '性能	',
        2,
        '',
        10,
        1637139540,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        159,
        20,
        1,
        8,
        '相机',
        2,
        '',
        10,
        1637139541,
        1
    );

INSERT INTO
    `goods_attr`
VALUES (
        160,
        20,
        1,
        9,
        '支持蓝牙',
        3,
        '是\r\n',
        10,
        1637139541,
        1
    );

-- ----------------------------
-- Table structure for goods_cate
-- ----------------------------
DROP TABLE IF EXISTS `goods_cate`;

CREATE TABLE `goods_cate` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `cate_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `template` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `pid` int(0) NULL DEFAULT NULL,
    `sub_title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `keywords` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `description` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `status` tinyint(1) NULL DEFAULT NULL,
    `sort` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `add_time` int(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_cate
-- ----------------------------
INSERT INTO
    `goods_cate`
VALUES (
        1,
        '手机',
        '',
        '',
        '',
        0,
        '手机',
        '手机',
        '手机',
        1,
        '10',
        1582461745
    );

INSERT INTO
    `goods_cate`
VALUES (
        2,
        '小米11 Pro pro',
        'static/upload/2020223/1582463294.png',
        '',
        'baidu/product/aaa.html',
        1,
        '小米10 Pro',
        '小米10 Pro',
        '小米10 Pro',
        1,
        '0',
        1582463294
    );

INSERT INTO
    `goods_cate`
VALUES (
        3,
        'Redmi 8',
        'static/upload/2020223/1582463357.png',
        'http://www.baidu.com',
        '11',
        1,
        'Redmi 8 11',
        'Redmi 8 111',
        'Redmi 8 111',
        1,
        '11',
        1582463357
    );

INSERT INTO
    `goods_cate`
VALUES (
        4,
        '电视 盒子',
        '',
        '',
        '',
        0,
        '电视 盒子',
        '电视 盒子',
        '电视 盒子',
        1,
        '0',
        1582463515
    );

INSERT INTO
    `goods_cate`
VALUES (
        5,
        '小米电视5 55英寸',
        'static/upload/2020223/1582464603.png',
        '',
        '',
        4,
        '小米电视5 55英寸',
        '小米电视5 55英寸',
        '小米电视5 55英寸',
        1,
        '0',
        1582464603
    );

INSERT INTO
    `goods_cate`
VALUES (
        6,
        '家电',
        '',
        '',
        '',
        0,
        '',
        '',
        '',
        1,
        '0',
        1582513219
    );

INSERT INTO
    `goods_cate`
VALUES (
        7,
        '出行 穿戴',
        '',
        '',
        '',
        0,
        '',
        '',
        '',
        1,
        '0',
        1582513235
    );

INSERT INTO
    `goods_cate`
VALUES (
        8,
        '智能 路由器',
        '',
        '',
        '',
        0,
        '',
        '',
        '',
        1,
        '0',
        1582513270
    );

INSERT INTO
    `goods_cate`
VALUES (
        9,
        '电源 配件',
        '',
        '',
        '',
        0,
        '',
        '',
        '',
        1,
        '0',
        1582513285
    );

INSERT INTO
    `goods_cate`
VALUES (
        10,
        '健康 儿童',
        'static/upload/20211028/1635413604527197900.jpg',
        '',
        '',
        0,
        '',
        '',
        '',
        1,
        '0',
        1582513300
    );

INSERT INTO
    `goods_cate`
VALUES (
        11,
        '耳机 音响',
        '',
        '',
        '',
        0,
        '',
        '',
        '',
        1,
        '0',
        1582513338
    );

INSERT INTO
    `goods_cate`
VALUES (
        12,
        '生活 箱包',
        '',
        '',
        '',
        0,
        '',
        '',
        '',
        1,
        '0',
        1582513349
    );

INSERT INTO
    `goods_cate`
VALUES (
        13,
        '冰箱',
        'static/upload/2020224/1582513945.jpg',
        '',
        '',
        6,
        '冰箱',
        '冰箱',
        '冰箱',
        1,
        '0',
        1582513945
    );

INSERT INTO
    `goods_cate`
VALUES (
        14,
        '微波炉',
        'static/upload/2020224/1582514001.jpg',
        '',
        '',
        6,
        '',
        '',
        '',
        1,
        '0',
        1582513960
    );

INSERT INTO
    `goods_cate`
VALUES (
        15,
        '小米手表',
        'static/upload/2020224/1582514113.png',
        '',
        '',
        7,
        '小米手表',
        '小米手表',
        '小米手表',
        1,
        '0',
        1582514113
    );

INSERT INTO
    `goods_cate`
VALUES (
        16,
        '平衡车',
        'static/upload/2020224/1582514151.jpg',
        '',
        '',
        7,
        '平衡车',
        '平衡车',
        '平衡车',
        1,
        '0',
        1582514151
    );

INSERT INTO
    `goods_cate`
VALUES (
        17,
        '路由器',
        'static/upload/2020224/1582514289.png',
        '',
        '',
        8,
        '路由器',
        '路由器',
        '路由器',
        1,
        '0',
        1582514289
    );

INSERT INTO
    `goods_cate`
VALUES (
        18,
        '摄像机',
        'static/upload/2020224/1582514318.jpg',
        '',
        '',
        8,
        '摄像机',
        '摄像机',
        '摄像机',
        1,
        '0',
        1582514318
    );

INSERT INTO
    `goods_cate`
VALUES (
        19,
        '全屏电视55寸',
        'static/upload/2020224/1582514664.jpg',
        '',
        '',
        4,
        '',
        '',
        '',
        1,
        '0',
        1582514664
    );

INSERT INTO
    `goods_cate`
VALUES (
        20,
        '移动电源',
        'static/upload/2020224/1582514810.png',
        '',
        '',
        9,
        '移动电源',
        '移动电源',
        '移动电源',
        1,
        '0',
        1582514810
    );

INSERT INTO
    `goods_cate`
VALUES (
        35,
        'Xiaomi 10S',
        'static/upload/20211102/1635841294026066400.png',
        '',
        '',
        1,
        '',
        '',
        '',
        1,
        '10',
        1635817714
    );

INSERT INTO
    `goods_cate`
VALUES (
        36,
        'Xiaomi Civi',
        'static/upload/20211102/1635841252665099500.png',
        '',
        '',
        1,
        '',
        '',
        '',
        1,
        '10',
        1635841252
    );

INSERT INTO
    `goods_cate`
VALUES (
        37,
        'Xiaomi MIX 4',
        'static/upload/20211102/1635841362004932300.png',
        '',
        '',
        1,
        '',
        '',
        '',
        1,
        '10',
        1635841362
    );

INSERT INTO
    `goods_cate`
VALUES (
        38,
        'Redmi K30S 至尊纪念版',
        'static/upload/20211102/1635841411131518300.png',
        '',
        '',
        1,
        '',
        '',
        '',
        1,
        '10',
        1635841411
    );

-- ----------------------------
-- Table structure for goods_color
-- ----------------------------
DROP TABLE IF EXISTS `goods_color`;

CREATE TABLE `goods_color` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `color_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `color_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `status` int(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_color
-- ----------------------------
INSERT INTO `goods_color` VALUES (1, '红色', 'red', 1);

INSERT INTO `goods_color` VALUES (2, '黑色', '#000', 1);

INSERT INTO `goods_color` VALUES (3, '黄色', 'yellow', 1);

INSERT INTO `goods_color` VALUES (4, '金色', '#ebf10f', 1);

INSERT INTO `goods_color` VALUES (5, '灰色', '#eee', 1);

INSERT INTO `goods_color` VALUES (6, '紫色', '#9932CD ', 1);

INSERT INTO `goods_color` VALUES (7, '淡绿色', '#90EE90', 1);

INSERT INTO `goods_color` VALUES (8, '蓝色', 'blue', NULL);

-- ----------------------------
-- Table structure for goods_image
-- ----------------------------
DROP TABLE IF EXISTS `goods_image`;

CREATE TABLE `goods_image` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `goods_id` int(0) NULL DEFAULT NULL,
    `img_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `color_id` int(0) NULL DEFAULT NULL,
    `sort` int(0) NULL DEFAULT NULL,
    `add_time` int(0) NULL DEFAULT NULL,
    `status` tinyint(1) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_image
-- ----------------------------
INSERT INTO
    `goods_image`
VALUES (
        3,
        36,
        'static/upload/20211009/1633755415645620800.png',
        3,
        10,
        1633755417,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        4,
        36,
        'static/upload/20211009/1633755415656163100.png',
        1,
        10,
        1633755417,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        5,
        37,
        'static/upload/20211009/1633755740718752300.png',
        2,
        10,
        1633755741,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        6,
        37,
        'static/upload/20211009/1633755740714630100.jpg',
        4,
        10,
        1633755741,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        7,
        38,
        'static/upload/20211009/1633755956051077200.png',
        0,
        10,
        1633755959,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        8,
        38,
        'static/upload/20211009/1633755956136482100.png',
        0,
        10,
        1633755959,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        9,
        38,
        'static/upload/20211009/1633755956135954600.jpg',
        0,
        10,
        1633755959,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        12,
        40,
        'static/upload/20211029/1635503037433844200.jpg',
        0,
        10,
        1635503038,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        13,
        40,
        'static/upload/20211029/1635503037587034300.jpg',
        0,
        10,
        1635503038,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        14,
        21,
        'static/upload/20211101/1635736448687849200.jpg',
        0,
        10,
        1635736455,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        19,
        43,
        'static/upload/20211102/1635841578192734800.png',
        0,
        10,
        1635841580,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        20,
        44,
        'static/upload/20211102/1635841907018281600.jpg',
        0,
        10,
        1635841908,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        21,
        19,
        'static/upload/20211116/1637052716813265400.jpg',
        0,
        10,
        1637052718,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        22,
        19,
        'static/upload/20211116/1637052716850583600.jpg',
        0,
        10,
        1637052718,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        23,
        19,
        'static/upload/20211116/1637052716884345200.jpg',
        0,
        10,
        1637052718,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        27,
        20,
        'static/upload/20211116/1637063586126852600.jpg',
        8,
        10,
        1637063587,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        28,
        53,
        'static/upload/20211116/1637063716307468700.jpg',
        2,
        10,
        1637063716,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        29,
        53,
        'static/upload/20211116/1637063716342727800.jpg',
        2,
        10,
        1637063716,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        30,
        53,
        'static/upload/20211117/1637138323678153500.jpg',
        6,
        10,
        1637138326,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        31,
        53,
        'static/upload/20211117/1637138323703994500.jpg',
        0,
        10,
        1637138326,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        32,
        53,
        'static/upload/20211117/1637138323728025100.jpg',
        7,
        10,
        1637138326,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        33,
        53,
        'static/upload/20211117/1637138324148951200.jpg',
        7,
        10,
        1637138326,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        34,
        53,
        'static/upload/20211117/1637138324173917100.jpg',
        7,
        10,
        1637138326,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        35,
        20,
        'static/upload/20211117/1637139106080100100.jpg',
        5,
        10,
        1637139108,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        36,
        20,
        'static/upload/20211117/1637139106095482200.jpg',
        5,
        10,
        1637139108,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        37,
        20,
        'static/upload/20211117/1637139106121304700.jpg',
        7,
        10,
        1637139108,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        38,
        20,
        'static/upload/20211117/1637139106142051100.jpg',
        2,
        10,
        1637139108,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        39,
        20,
        'static/upload/20211117/1637139106176296000.jpg',
        8,
        10,
        1637139108,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        40,
        20,
        'static/upload/20211117/1637139106213056300.jpg',
        7,
        10,
        1637139108,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        41,
        54,
        'static/upload/20211117/1637139499477843800.jpg',
        2,
        10,
        1637139501,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        42,
        54,
        'static/upload/20211117/1637139499510094900.jpg',
        2,
        10,
        1637139501,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        43,
        54,
        'static/upload/20211117/1637139499528738400.jpg',
        8,
        10,
        1637139501,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        44,
        54,
        'static/upload/20211117/1637139499755592600.jpg',
        8,
        10,
        1637139501,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        45,
        50,
        'static/upload/20211118/1637202369917645700.jpg',
        0,
        10,
        1637202371,
        1
    );

INSERT INTO
    `goods_image`
VALUES (
        46,
        50,
        'static/upload/20211118/1637202369935282300.jpg',
        0,
        10,
        1637202371,
        1
    );

-- ----------------------------
-- Table structure for goods_type
-- ----------------------------
DROP TABLE IF EXISTS `goods_type`;

CREATE TABLE `goods_type` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `status` int(0) NULL DEFAULT NULL,
    `add_time` int(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_type
-- ----------------------------
INSERT INTO `goods_type` VALUES ( 1, '手机', '手机22', 1, 1632299505 );

INSERT INTO `goods_type` VALUES (2, '电脑', '电脑', 0, 1632299512);

INSERT INTO `goods_type` VALUES ( 3, '笔记本', '笔记本', 1, 1632299526 );

INSERT INTO `goods_type` VALUES ( 4, '路由器', '路由器', 0, 1632299535 );

INSERT INTO `goods_type` VALUES (9, '衣服', '衣服', 1, 1632361292);

-- ----------------------------
-- Table structure for goods_type_attribute
-- ----------------------------
DROP TABLE IF EXISTS `goods_type_attribute`;

CREATE TABLE `goods_type_attribute` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `cate_id` int(0) NULL DEFAULT NULL,
    `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `attr_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `attr_value` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `status` tinyint(1) NULL DEFAULT NULL,
    `sort` int(0) NULL DEFAULT NULL,
    `add_time` int(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `cate_id` (`cate_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of goods_type_attribute
-- ----------------------------
INSERT INTO
    `goods_type_attribute`
VALUES (
        1,
        1,
        '基本信息',
        '1',
        '',
        1,
        10,
        1632299512
    );

INSERT INTO
    `goods_type_attribute`
VALUES (
        2,
        2,
        '主体',
        '3',
        '111\r\n1111',
        1,
        19,
        1632299512
    );

INSERT INTO
    `goods_type_attribute`
VALUES (
        3,
        2,
        '内存',
        '1',
        '',
        1,
        NULL,
        1632299512
    );

INSERT INTO
    `goods_type_attribute`
VALUES (
        4,
        2,
        '硬盘',
        '1',
        '',
        1,
        NULL,
        1632299512
    );

INSERT INTO
    `goods_type_attribute`
VALUES (
        5,
        2,
        '显示器',
        '1',
        '',
        1,
        111,
        1582361804
    );

INSERT INTO
    `goods_type_attribute`
VALUES (
        6,
        2,
        '支持蓝牙',
        '3',
        '是\r\n否',
        1,
        1011,
        1582362691
    );

INSERT INTO
    `goods_type_attribute`
VALUES (
        7,
        1,
        '性能	',
        '2',
        '',
        1,
        111,
        1632299512
    );

INSERT INTO
    `goods_type_attribute`
VALUES (
        8,
        1,
        '相机',
        '2',
        '',
        1,
        0,
        1632299512
    );

INSERT INTO
    `goods_type_attribute`
VALUES (
        9,
        1,
        '支持蓝牙',
        '3',
        '是\r\n否',
        1,
        0,
        1591844649
    );

INSERT INTO
    `goods_type_attribute`
VALUES (
        10,
        4,
        '是否支持蓝牙',
        '3',
        '是\r\n否',
        1,
        1022,
        1632370943
    );

INSERT INTO
    `goods_type_attribute`
VALUES (
        12,
        3,
        '尺寸1',
        '1',
        '',
        1,
        10,
        1632388221
    );

-- ----------------------------
-- Table structure for manager
-- ----------------------------
DROP TABLE IF EXISTS `manager`;

CREATE TABLE `manager` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `mobile` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `status` tinyint(1) NULL DEFAULT NULL,
    `role_id` int(0) NULL DEFAULT NULL,
    `add_time` int(0) NULL DEFAULT NULL,
    `is_super` tinyint(1) NULL DEFAULT 0,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `role_id` (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of manager
-- ----------------------------
INSERT INTO
    `manager`
VALUES (
        1,
        'admin',
        'e10adc3949ba59abbe56e057f20f883e',
        '152016111',
        '5188611114@qq.com',
        1,
        9,
        0,
        1
    );

INSERT INTO
    `manager`
VALUES (
        2,
        'zhangsan',
        'e10adc3949ba59abbe56e057f20f883e',
        '1520111122',
        '342338691122@qq.com',
        1,
        14,
        1581661532,
        0
    );

INSERT INTO
    `manager`
VALUES (
        6,
        'lisi',
        'e10adc3949ba59abbe56e057f20f883e',
        '1520171111',
        '11114292@qq.com',
        1,
        16,
        1631156378,
        0
    );

-- ----------------------------
-- Table structure for nav
-- ----------------------------
DROP TABLE IF EXISTS `nav`;

CREATE TABLE `nav` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `position` tinyint(1) NULL DEFAULT NULL,
    `is_opennew` tinyint(1) NULL DEFAULT NULL,
    `relation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `sort` int(0) NULL DEFAULT NULL,
    `status` tinyint(1) NULL DEFAULT NULL,
    `add_time` int(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of nav
-- ----------------------------
INSERT INTO
    `nav`
VALUES (
        1,
        '小米商城',
        'http://www.baidu.com',
        2,
        2,
        '21,22,23,24',
        10,
        1,
        1592919226
    );

INSERT INTO
    `nav`
VALUES (
        2,
        'MIUI',
        'http://www.baidu.com',
        1,
        1,
        '1',
        10,
        1,
        1592921999
    );

INSERT INTO
    `nav`
VALUES (
        3,
        '小米手机',
        'https://shouji.mi.com/',
        2,
        2,
        '19,20',
        10,
        1,
        1592922081
    );

INSERT INTO
    `nav`
VALUES (
        4,
        '小米电视',
        'https://ds.mi.com/',
        2,
        2,
        '23,24',
        10,
        1,
        1592922273
    );

INSERT INTO
    `nav`
VALUES (
        5,
        '路由器',
        'http://bbs.baidu.com',
        2,
        1,
        '25',
        10,
        1,
        1592922331
    );

INSERT INTO
    `nav`
VALUES (
        8,
        '云服务',
        'https://i.mi.com/',
        1,
        2,
        '2',
        10,
        1,
        1593529309
    );

INSERT INTO
    `nav`
VALUES (
        9,
        '金融',
        'https://jr.mi.com/?from=micom',
        1,
        1,
        '1',
        10,
        1,
        1593529329
    );

INSERT INTO
    `nav`
VALUES (
        10,
        '有品',
        'https://youpin.mi.com/',
        1,
        1,
        '1',
        10,
        1,
        1593529346
    );

INSERT INTO
    `nav`
VALUES (
        11,
        '家电',
        '',
        2,
        1,
        '1',
        10,
        1,
        1593529451
    );

INSERT INTO
    `nav`
VALUES (
        12,
        '智能电视',
        '',
        2,
        1,
        '37',
        10,
        1,
        1593529470
    );

INSERT INTO
    `nav`
VALUES (
        14,
        '小米帮助中心2',
        'http://www.baidu.com',
        3,
        2,
        '12,13,14',
        10,
        1,
        1634788777
    );

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;

CREATE TABLE `role` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `status` tinyint(1) NULL DEFAULT NULL,
    `add_time` int(0) NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ( 9, '超级管理员', '我是一个超级管理员', 1, 1631072961 );

INSERT INTO `role` VALUES ( 14, '软件部门', '软件部门', 1, 1631075350 );

INSERT INTO `role` VALUES ( 16, '销售部门', '销售部门', 1, 1631589828 );

-- ----------------------------
-- Table structure for role_access
-- ----------------------------
DROP TABLE IF EXISTS `role_access`;

CREATE TABLE `role_access` (
    `role_id` int(0) NOT NULL,
    `access_id` int(0) NOT NULL,
    INDEX `role_id` (`role_id`) USING BTREE,
    INDEX `access_id` (`access_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_access
-- ----------------------------
INSERT INTO `role_access` VALUES (14, 52);

INSERT INTO `role_access` VALUES (14, 54);

INSERT INTO `role_access` VALUES (14, 55);

INSERT INTO `role_access` VALUES (14, 56);

INSERT INTO `role_access` VALUES (14, 57);

INSERT INTO `role_access` VALUES (14, 76);

INSERT INTO `role_access` VALUES (14, 53);

INSERT INTO `role_access` VALUES (14, 59);

INSERT INTO `role_access` VALUES (14, 60);

INSERT INTO `role_access` VALUES (14, 61);

INSERT INTO `role_access` VALUES (14, 62);

INSERT INTO `role_access` VALUES (14, 78);

INSERT INTO `role_access` VALUES (14, 79);

INSERT INTO `role_access` VALUES (14, 80);

INSERT INTO `role_access` VALUES (14, 81);

INSERT INTO `role_access` VALUES (9, 52);

INSERT INTO `role_access` VALUES (9, 54);

INSERT INTO `role_access` VALUES (9, 55);

INSERT INTO `role_access` VALUES (9, 53);

INSERT INTO `role_access` VALUES (9, 59);

INSERT INTO `role_access` VALUES (9, 60);

INSERT INTO `role_access` VALUES (9, 61);

INSERT INTO `role_access` VALUES (9, 62);

INSERT INTO `role_access` VALUES (9, 63);

INSERT INTO `role_access` VALUES (9, 64);

INSERT INTO `role_access` VALUES (9, 67);

INSERT INTO `role_access` VALUES (9, 82);

INSERT INTO `role_access` VALUES (9, 83);

INSERT INTO `role_access` VALUES (9, 84);

INSERT INTO `role_access` VALUES (9, 85);

INSERT INTO `role_access` VALUES (9, 70);

INSERT INTO `role_access` VALUES (9, 71);

INSERT INTO `role_access` VALUES (16, 53);

INSERT INTO `role_access` VALUES (16, 59);

INSERT INTO `role_access` VALUES (16, 60);

INSERT INTO `role_access` VALUES (16, 61);

INSERT INTO `role_access` VALUES (16, 62);

INSERT INTO `role_access` VALUES (16, 78);

INSERT INTO `role_access` VALUES (16, 79);

INSERT INTO `role_access` VALUES (16, 80);

INSERT INTO `role_access` VALUES (16, 81);

INSERT INTO `role_access` VALUES (16, 63);

INSERT INTO `role_access` VALUES (16, 64);

INSERT INTO `role_access` VALUES (16, 67);

INSERT INTO `role_access` VALUES (16, 82);

INSERT INTO `role_access` VALUES (16, 83);

INSERT INTO `role_access` VALUES (16, 84);

INSERT INTO `role_access` VALUES (16, 85);

-- ----------------------------
-- Table structure for setting
-- ----------------------------
DROP TABLE IF EXISTS `setting`;

CREATE TABLE `setting` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    `site_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `site_logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `site_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `site_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `no_picture` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `site_icp` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `site_tel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `search_keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `tongji_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `appid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `app_secret` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `end_point` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `bucket_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `oss_status` tinyint(1) NULL DEFAULT NULL,
    `oss_domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `thumbnail_size` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of setting
-- ----------------------------
INSERT INTO
    `setting`
VALUES (
        1,
        'Gin仿小米商城项目3333333333',
        'static/upload/20211029/1635492623215921000.jpg',
        '小米',
        '222222222',
        'static/upload/20211029/1635472950579882800.jpg',
        '2422',
        '24',
        '24',
        '11111',
        'GJoqWHXB2c9S9gwP',
        'Lgf3weXuWITUUb17vDJfveg1jmKEe9',
        'oss-cn-beijing.aliyuncs.com',
        'baidu',
        0,
        'http://bee.apiying.com/',
        '100,200,400'
    );

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
    `id` int(0) NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
