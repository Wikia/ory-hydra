-- Migration generated by the command below; DO NOT EDIT.
-- hydra:generate hydra migrate gen

UPDATE hydra_oauth2_trusted_jwt_bearer_issuer SET nid = (SELECT id FROM networks LIMIT 1);
ALTER TABLE hydra_oauth2_trusted_jwt_bearer_issuer MODIFY `nid` char(36) NOT NULL;
