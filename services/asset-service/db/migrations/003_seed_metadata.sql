-- db/migrations/003_seed_metadata.sql

-- 1. Seed Asset Classes
INSERT INTO cis_asset_class (name, definition) VALUES
('Devices', 'Hardware assets: end-user devices, servers, IoT, and mobile.'),
('Software', 'Operating systems, applications, and specialized software.'),
('Data', 'Information assets: sensitive, log, and operational data.'),
('Users', 'Workforce accounts, service accounts, and administrators.'),
('Network', 'Network infrastructure: routers, switches, and firewalls.'),
('Documentation', 'Policies, procedures, and security diagrams.')
ON CONFLICT (name) DO NOTHING;

-- 2. SEED DEVICE TYPES (The missing link!)
-- This maps specific types to the classes created above.
INSERT INTO device_type (asset_class_id, name) VALUES
((SELECT id FROM cis_asset_class WHERE name = 'Devices'), 'Laptop'),
((SELECT id FROM cis_asset_class WHERE name = 'Devices'), 'Server'),
((SELECT id FROM cis_asset_class WHERE name = 'Devices'), 'Workstation'),
((SELECT id FROM cis_asset_class WHERE name = 'Devices'), 'Virtual Machine'),
((SELECT id FROM cis_asset_class WHERE name = 'Network'), 'Switch'),
((SELECT id FROM cis_asset_class WHERE name = 'Network'), 'Router'),
((SELECT id FROM cis_asset_class WHERE name = 'Network'), 'Firewall'),
((SELECT id FROM cis_asset_class WHERE name = 'Network'), 'Access Point')
ON CONFLICT (asset_class_id, name) DO NOTHING;

-- 3. Seed Criticality
INSERT INTO cis_asset_criticality (name, value) VALUES
('Low', 1), ('Medium', 2), ('High', 3), ('Critical', 4)
ON CONFLICT (name) DO NOTHING;

-- 4. Seed Asset Sources
INSERT INTO asset_source (name) VALUES
('CrowdStrike'), ('Qualys'), ('SolarWinds'), ('vCenter'), ('Active Directory'), ('Manual')
ON CONFLICT (name) DO NOTHING;

-- 5. Seed Issue Sources
INSERT INTO issue_source (name) VALUES
('Tenable'), ('Rapid7'), ('Cisco Advisory'), ('GitHub Dependabot'), ('Manual Audit'), ('NIST NVD')
ON CONFLICT (name) DO NOTHING;

-- 6. Seed Issue Statuses
INSERT INTO cis_issue_status (name) VALUES
('New'), ('Confirmed'), ('In Progress'), ('Remediated'), ('Resolved'), ('False Positive'), ('Risk Accepted')
ON CONFLICT (name) DO NOTHING;

-- 7. Seed Issue Categories
INSERT INTO cis_issue_category (name, description) VALUES
('Configuration Drift', 'A setting that was previously secure but has changed over time.'),
('Vulnerability', 'A known flaw in software or hardware (often tracked via CVEs).'),
('Non-Compliance', 'An asset that fails to meet a specific CIS Benchmark recommendation.'),
('Unauthorized Asset', 'An asset found on the network that is not in the authorized inventory.'),
('Account Issue', 'Problems related to user identity, privileges, or dormant accounts.')
ON CONFLICT (name) DO NOTHING;

---- create above / drop below ----

TRUNCATE cis_asset_class, cis_asset_criticality, asset_source, issue_source, cis_issue_status, cis_issue_category, device_type CASCADE;
