-- Migration generated by the command below; DO NOT EDIT.
-- hydra:generate hydra migrate gen


ALTER TABLE hydra_oauth2_flow DROP CONSTRAINT hydra_oauth2_flow_chk;

ALTER TABLE hydra_oauth2_flow DROP COLUMN requested_scope;
ALTER TABLE hydra_oauth2_flow DROP COLUMN requested_at_audience;
ALTER TABLE hydra_oauth2_flow DROP COLUMN amr;
ALTER TABLE hydra_oauth2_flow DROP COLUMN granted_scope;
ALTER TABLE hydra_oauth2_flow DROP COLUMN granted_at_audience;
