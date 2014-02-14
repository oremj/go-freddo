%define  debug_package %{nil}
%global __strip /bin/true
%global rev ad8bd97c2b3923bd928d4e6afe096df7b55273c6
%global shortrev %(r=%{rev}; echo ${r:0:12})

Name: go-freddo
Version: 0.0.1
Release: r1.%{shortrev}%{?dist}
Summary: A remote process runner.

License: MPL
URL: https://github.com/oremj/go-freddo
Source0: go-freddo-%{version}.tar.gz
BuildRoot: %(mktemp -ud %{_tmppath}/%{name}-%{version}-%{release}-XXXXXX)


%description
A remote process runner.

%prep
%setup -q -n go-freddo-master

%build
./build


%install
rm -rf %{buildroot}
install -d %{buildroot}%{_bindir}
install ./go-freddo %{buildroot}%{_bindir}/go-freddo


%clean
rm -rf %{buildroot}


%files
%defattr(-,root,root,-)
%{_bindir}/go-freddo

%changelog
