m = require("mymodule")
--m.myfunc()
print(m.nowTime)
print("Случайное число от 100 до 1000:", m.randNumber)
a, b = m.rn1 , m.rn2
res = mult(a, b)
print(string.format("%d x %d = %d", a, b, res))
d = data()
getInLua(res*d)

print("Table")
t = {a = 10, b = 20, c = 30}
s = {a = 1, c = 3}

print("Table t:")
for key, val in pairs(t) do
  print(key, val)
  if s[key] == nil then
     s[key] = 0
  end
end

print("Table s:")
for key, val in pairs(s) do
  print(key, val)
end

print("Operation:")
for key, val in pairs(t) do
  print('t.' .. key, ' = ' .. val)
  print('s.' .. key, ' = ' .. s[key])
end

isp = {
'DigitalOcean LLC',
'OVH SAS',
'Opera Software AS',
'NetArt Group s.r.o.',
'Online S.A.S.',
'My ISP SARL',
'M247 Europe SRL',
'ServiHosting Networks S.L.',
'WorldStream B.V.',
'RunAbove',
'Linode LLC',
'Hetzner Online AG',
'Worldstream Latam B.V',
'Lancom Ltd.',
'Link Data Group',
'Google LLC',
'OVH Hosting Inc.',
'AltusHost B.V.',
'Contabo GmbH',
'Net Tech Ltd',
'M247 Ltd',
'Fasthosts Internet Limited',
'Digital Ocean Inc.',
'FDCServers.net',
'DataCamp s.r.o.',
'PPMAN Services Srl',
'Vultr Holdings LLC',
'CloudIP LLC',
'Micfo LLC.',
'Greenhost BV',
'IPv4 Management SRL',
'DataCamp Limited',
'Choopa LLC',
'G-Core Labs S.A.',
'Madgenius.com',
'Secure Data Systems SRL',
'IFX Networks Colombia',
'Host Europe GmbH',
'Subnet LLC',
'Arsys Internet S.L.',
'UpCloud Ltd',
'Frantech Solutions',
'LLC TRC Fiord',
'Hosting Services Inc.',
'GlobalTelehost Corp.',
'BuyVM',
'Linode',
'EGIHosting',
'Triple C Cloud Computing Ltd.',
'Langate Ltd',
'UK Web.Solutions Direct Ltd',
'LeaseWeb Deutschland GmbH',
'SysEleven GmbH',
'Zscaler Inc.',
'Online SAS',
'Versaweb LLC',
'Amazon Technologies Inc.',
'Dedibox SAS',
'CloudRoute LLC',
'New York City Cloud',
'QuadraNet Enterprises LLC',
'LeaseWeb Netherlands B.V.',
'UpCloud USA Inc',
'iomart Hosting Limited',
'Opera Software Americas LLC',
'Total Server Solutions L.L.C.',
'KVCHosting.com LLC',
'Colocation America Corporation',
'WorldStream LATAM B.V',
'ATOMOHOST LLC',
'Amazon.com Inc.',
'2 Cloud Ltd.',
'Amazon Data Services UK',
'OVH US LLC',
'Host Wagon LLC',
'Hosting and Colocation Services',
'LeaseWeb USA Inc.',
'LeaseWeb UK Limited',
'The Cloud Networks Limited',
'DataCheap Ltd.',
'Microsoft Corporation',
}

print("isp :")
str = ""
for key, val in pairs(isp) do
  print(key, val)
  if key == 1 then
     str = str .. "'" .. val .. "'"
  end
  str = str .. ", '" .. val .. "'"
end

print("str :", str)

con = {
'COM',
'EDU',
'MOB',
'GOV',
'ORG',
'ISP/MOB',
'ISP',
'SES',
'DCH',
}
