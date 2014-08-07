%define  debug_package %{nil}
%global __strip /bin/true

Name: go-freddo
Version: 1.0.1
Release: r1.%{?dist}
Summary: A remote process runner.

License: MPL
URL: https://github.com/oremj/go-freddo
Source0: go-freddo-%{version}.tar.gz
BuildRoot: %(mktemp -ud %{_tmppath}/%{name}-%{version}-%{release}-XXXXXX)


%description
A remote process runner.

%prep
%setup -q -n go-freddo-%{version}

%build
./build

%install
rm -rf %{buildroot}
install -d %{buildroot}%{_bindir}
install ./go-freddo %{buildroot}%{_bindir}/freddo


%clean
rm -rf %{buildroot}


%files
%defattr(-,root,root,-)
%{_bindir}/freddo

%changelog
