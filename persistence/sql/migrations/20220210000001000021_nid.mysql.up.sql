-- Migration generated by the command below; DO NOT EDIT.
-- hydra:generate hydra migrate gen

ALTER TABLE `hydra_oauth2_jti_blacklist` ADD COLUMN `nid` char(36);
ALTER TABLE `hydra_oauth2_jti_blacklist` ADD CONSTRAINT `hydra_oauth2_jti_blacklist_nid_fk_idx` FOREIGN KEY (`nid`) REFERENCES `networks` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE;
