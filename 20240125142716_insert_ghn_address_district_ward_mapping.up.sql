
BEGIN;

SET FOREIGN_KEY_CHECKS = 0;

DELETE FROM shipping_provider_district_mappings WHERE shipping_provider_code = 'GHN'
DELETE FROM shipping_provider_ward_mappings WHERE shipping_provider_code = 'GHN'

SET FOREIGN_KEY_CHECKS = 1;

COMMIT;