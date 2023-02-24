/*
 Navicat Premium Data Transfer

 Source Server         : root
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : tik_tok

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 12/02/2023 20:06:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_comment_videos
-- ----------------------------
DROP TABLE IF EXISTS `user_comment_videos`;
CREATE TABLE `user_comment_videos`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `video_id` bigint NULL DEFAULT NULL COMMENT '视频id',
  `comment_text` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '评论内容',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_comment_videos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_comment_videos
-- ----------------------------
INSERT INTO `user_comment_videos` VALUES (1, '2022-12-14 01:11:27', '2022-12-14 01:11:27', NULL, 3, 2, '444');

-- ----------------------------
-- Table structure for user_focus_ons
-- ----------------------------
DROP TABLE IF EXISTS `user_focus_ons`;
CREATE TABLE `user_focus_ons`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `to_user_id` bigint NULL DEFAULT NULL COMMENT '对方用户id',
  `is_follow` tinyint(1) NULL DEFAULT NULL COMMENT '关注状态:true-已关注，false-未关注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_focus_ons_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_focus_ons
-- ----------------------------
INSERT INTO `user_focus_ons` VALUES (1, NULL, NULL, NULL, 3, 3, 0);

-- ----------------------------
-- Table structure for user_like_videos
-- ----------------------------
DROP TABLE IF EXISTS `user_like_videos`;
CREATE TABLE `user_like_videos`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `video_id` bigint NULL DEFAULT NULL COMMENT '视频id',
  `is_favorite` tinyint(1) NULL DEFAULT NULL COMMENT '点赞状态',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_like_videos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_like_videos
-- ----------------------------
INSERT INTO `user_like_videos` VALUES (1, '2022-12-14 01:11:15', '2022-12-14 01:11:20', NULL, 3, 2, 1);

-- ----------------------------
-- Table structure for user_release_videos
-- ----------------------------
DROP TABLE IF EXISTS `user_release_videos`;
CREATE TABLE `user_release_videos`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `video_id` bigint NULL DEFAULT NULL COMMENT '视频id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_release_videos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_release_videos
-- ----------------------------
INSERT INTO `user_release_videos` VALUES (1, '2022-12-08 00:40:59', NULL, NULL, 3, 2);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户登录密码',
  `follow_count` bigint NULL DEFAULT NULL COMMENT '关注总数',
  `follower_count` bigint NULL DEFAULT NULL COMMENT '粉丝总数',
  `like_total` bigint NULL DEFAULT NULL COMMENT '点赞总数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2022-12-05 23:54:44', '2022-12-06 23:54:50', '2022-12-01 00:10:27', 'hk', '123456', 11, 11, 11);
INSERT INTO `users` VALUES (2, '2022-12-01 00:11:21', NULL, NULL, 'hk2', '123456', 22, 2, 22);
INSERT INTO `users` VALUES (3, '2022-12-14 00:16:32', '2022-12-14 00:16:32', NULL, 'hk3', 'e10adc3949ba59abbe56e057f20f883e', 0, 0, 0);
INSERT INTO `users` VALUES (4, '2022-12-14 01:17:51', '2022-12-14 01:17:51', NULL, '123', 'a8698009bce6d1b8c2128eddefc25aad', 0, 0, 0);

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频标题',
  `play_url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频播放地址',
  `cover_url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频封面地址',
  `favorite_count` bigint NULL DEFAULT NULL COMMENT '视频的点赞总数',
  `comment_count` bigint NULL DEFAULT NULL COMMENT '视频的评论总数',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文件名',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_videos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (1, '2022-12-12 23:52:06', '2022-12-13 23:52:13', '2022-11-15 00:04:49', 'test', 'http://101.43.131.145:81/vd/1.mp4', NULL, 1, 0, 'hk', '1', '1');
INSERT INTO `videos` VALUES (2, '2022-12-08 00:35:07', '2022-12-14 01:11:27', NULL, 'test1', 'http://101.43.131.145:81/vd/3.mp4', NULL, 3, 1, 'hk3', '1', '1');

SET FOREIGN_KEY_CHECKS = 1;
