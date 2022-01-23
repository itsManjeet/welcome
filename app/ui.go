package app

const UI = `<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated with glade 3.36.0 -->
<interface>
  <requires lib="gtk+" version="3.24"/>
  <object class="GtkApplicationWindow" id="mainWindow">
    <property name="can_focus">False</property>
    <property name="resizable">False</property>
    <property name="window_position">center-always</property>
    <property name="default_width">800</property>
    <property name="default_height">600</property>
    <property name="skip_taskbar_hint">True</property>
    <property name="skip_pager_hint">True</property>
    <property name="deletable">False</property>
    <child>
      <object class="GtkStack" id="stack">
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="halign">center</property>
            <property name="valign">center</property>
            <property name="margin_bottom">150</property>
            <property name="orientation">vertical</property>
            <child>
              <object class="GtkImage">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">19</property>
                <property name="margin_bottom">20</property>
                <property name="pixel_size">128</property>
                <property name="icon_name">rlxos</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="label" translatable="yes">Welcome!</property>
                <attributes>
                  <attribute name="weight" value="heavy"/>
                  <attribute name="scale" value="1.8"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">1</property>
                <property name="label" translatable="yes">Thanks for choosing rlxos GNU/Linux</property>
                <attributes>
                  <attribute name="weight" value="medium"/>
                  <attribute name="scale" value="1.2"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="name">page0</property>
            <property name="title" translatable="yes">page0</property>
          </packing>
        </child>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="orientation">vertical</property>
            <child>
              <object class="GtkImage">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">12</property>
                <property name="margin_bottom">5</property>
                <property name="pixel_size">64</property>
                <property name="icon_name">software-update-available-symbolic</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">2</property>
                <property name="label" translatable="yes">Whats New!</property>
                <attributes>
                  <attribute name="weight" value="heavy"/>
                  <attribute name="scale" value="1.5"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">2</property>
                <property name="label" translatable="yes">Build 2200 is preloaded with the features we promised in 2100</property>
                <attributes>
                  <attribute name="weight" value="normal"/>
                  <attribute name="scale" value="1.2"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkGrid">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_start">25</property>
                <property name="margin_end">25</property>
                <property name="margin_top">25</property>
                <property name="margin_bottom">25</property>
                <property name="row_spacing">25</property>
                <property name="column_spacing">25</property>
                <property name="row_homogeneous">True</property>
                <property name="column_homogeneous">True</property>
                <child>
                  <object class="GtkButton">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="valign">center</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Stability</property>
                            <attributes>
                              <attribute name="weight" value="semibold"/>
                              <attribute name="scale" value="1.05"/>
                            </attributes>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">First stable release of rlxos that can 
be installed on real hardware. 
However it still have some issues</property>
                            <property name="wrap">True</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="left_attach">0</property>
                    <property name="top_attach">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkButton">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="valign">center</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">PKGUPD</property>
                            <attributes>
                              <attribute name="weight" value="semibold"/>
                              <attribute name="scale" value="1.05"/>
                            </attributes>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Migration to every new PKGUPD
package management tool to install
packages that can't be distributed as
AppImage</property>
                            <property name="wrap">True</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="left_attach">1</property>
                    <property name="top_attach">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkButton">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="valign">center</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Refined Concepts</property>
                            <attributes>
                              <attribute name="weight" value="semibold"/>
                              <attribute name="scale" value="1.05"/>
                            </attributes>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Just like appimages, you can now 
update or replace your system just
by replacing the system image</property>
                            <property name="wrap">True</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="left_attach">2</property>
                    <property name="top_attach">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkButton">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="valign">center</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Source (src)</property>
                            <attributes>
                              <attribute name="weight" value="semibold"/>
                              <attribute name="scale" value="1.05"/>
                            </attributes>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Yaa! our homegrown programming 
language that helps your to 
automate your daily task
(experimental)</property>
                            <property name="wrap">True</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="left_attach">0</property>
                    <property name="top_attach">1</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkButton">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="valign">center</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Bolt</property>
                            <attributes>
                              <attribute name="weight" value="semibold"/>
                              <attribute name="scale" value="1.05"/>
                            </attributes>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">A untrained virtual assistant to help
you out in your daily tasks, train it
and use it.</property>
                            <property name="wrap">True</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="left_attach">1</property>
                    <property name="top_attach">1</property>
                  </packing>
                </child>
                <child>
                  <placeholder/>
                </child>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">3</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="name">page1</property>
            <property name="title" translatable="yes">page1</property>
            <property name="position">1</property>
          </packing>
        </child>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="orientation">vertical</property>
            <child>
              <object class="GtkImage">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">12</property>
                <property name="margin_bottom">5</property>
                <property name="pixel_size">64</property>
                <property name="icon_name">user-desktop-symbolic</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">2</property>
                <property name="label" translatable="yes">User Interface!</property>
                <attributes>
                  <attribute name="weight" value="heavy"/>
                  <attribute name="scale" value="1.5"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">2</property>
                <property name="label" translatable="yes">We try our best to provide a simple and user friendly interface.</property>
                <attributes>
                  <attribute name="weight" value="normal"/>
                  <attribute name="scale" value="1.2"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkImage">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">9</property>
                <property name="stock">gtk-missing-image</property>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">3</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="name">page2</property>
            <property name="title" translatable="yes">page2</property>
            <property name="position">2</property>
          </packing>
        </child>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="orientation">vertical</property>
            <child>
              <object class="GtkImage">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">12</property>
                <property name="margin_bottom">5</property>
                <property name="pixel_size">64</property>
                <property name="icon_name">gnome-software-symbolic</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">2</property>
                <property name="label" translatable="yes">Apps</property>
                <attributes>
                  <attribute name="weight" value="heavy"/>
                  <attribute name="scale" value="1.5"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">2</property>
                <property name="label" translatable="yes">You can get your applications from PKGUPD, AppImages and flatpaks</property>
                <attributes>
                  <attribute name="weight" value="normal"/>
                  <attribute name="scale" value="1.2"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkButtonBox">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_bottom">150</property>
                <property name="layout_style">spread</property>
                <child>
                  <object class="GtkButton">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <signal name="clicked" handler="onBazaarClicked" swapped="no"/>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="margin_start">10</property>
                        <property name="margin_end">10</property>
                        <property name="margin_top">25</property>
                        <property name="margin_bottom">25</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkImage">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="pixel_size">64</property>
                            <property name="icon_name">gnome-software</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Bazaar</property>
                            <attributes>
                              <attribute name="weight" value="semibold"/>
                              <attribute name="scale" value="1.1000000000000001"/>
                            </attributes>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Native packages and AppImages
provided by rlxos</property>
                            <property name="justify">center</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">2</property>
                          </packing>
                        </child>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">True</property>
                    <property name="fill">True</property>
                    <property name="position">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkButton">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <signal name="clicked" handler="onAppImageHubClicked" swapped="no"/>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="margin_left">10</property>
                        <property name="margin_right">10</property>
                        <property name="margin_start">10</property>
                        <property name="margin_end">10</property>
                        <property name="margin_top">25</property>
                        <property name="margin_bottom">25</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkImage">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="pixel_size">64</property>
                            <property name="icon_name">appimages</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">AppImageHub</property>
                            <attributes>
                              <attribute name="weight" value="semibold"/>
                              <attribute name="scale" value="1.1000000000000001"/>
                            </attributes>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Global Market of AppImages</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">2</property>
                          </packing>
                        </child>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">True</property>
                    <property name="fill">True</property>
                    <property name="position">1</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkButton">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <signal name="clicked" handler="onFlathubClicked" swapped="no"/>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="margin_left">10</property>
                        <property name="margin_right">10</property>
                        <property name="margin_start">10</property>
                        <property name="margin_end">10</property>
                        <property name="margin_top">25</property>
                        <property name="margin_bottom">25</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkImage">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="pixel_size">64</property>
                            <property name="icon_name">flatpak</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Flathub</property>
                            <attributes>
                              <attribute name="weight" value="semibold"/>
                              <attribute name="scale" value="1.1000000000000001"/>
                            </attributes>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Global market of
flatpak packages</property>
                            <property name="justify">center</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">2</property>
                          </packing>
                        </child>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">True</property>
                    <property name="fill">True</property>
                    <property name="position">2</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">3</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="name">page3</property>
            <property name="title" translatable="yes">page3</property>
            <property name="position">3</property>
          </packing>
        </child>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="orientation">vertical</property>
            <child>
              <object class="GtkImage">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">12</property>
                <property name="margin_bottom">5</property>
                <property name="pixel_size">64</property>
                <property name="icon_name">help-symbolic</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">2</property>
                <property name="label" translatable="yes">Help and Support</property>
                <attributes>
                  <attribute name="weight" value="heavy"/>
                  <attribute name="scale" value="1.5"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">2</property>
                <property name="label" translatable="yes">Feel free to join our little family</property>
                <attributes>
                  <attribute name="weight" value="normal"/>
                  <attribute name="scale" value="1.2"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkButtonBox">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_bottom">150</property>
                <property name="layout_style">spread</property>
                <child>
                  <object class="GtkButton">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <signal name="clicked" handler="onDocsClicked" swapped="no"/>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="valign">center</property>
                        <property name="margin_left">10</property>
                        <property name="margin_right">10</property>
                        <property name="margin_start">10</property>
                        <property name="margin_end">10</property>
                        <property name="margin_top">25</property>
                        <property name="margin_bottom">25</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkImage">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="pixel_size">64</property>
                            <property name="icon_name">help</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Documentations</property>
                            <attributes>
                              <attribute name="weight" value="semibold"/>
                              <attribute name="scale" value="1.1000000000000001"/>
                            </attributes>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Comprehensive documentation, 
which consists of a rlxos 
usage guide</property>
                            <property name="justify">center</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">2</property>
                          </packing>
                        </child>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">True</property>
                    <property name="fill">True</property>
                    <property name="position">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkButton">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <signal name="clicked" handler="onCommunityClicked" swapped="no"/>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="valign">center</property>
                        <property name="margin_left">10</property>
                        <property name="margin_right">10</property>
                        <property name="margin_start">10</property>
                        <property name="margin_end">10</property>
                        <property name="margin_top">25</property>
                        <property name="margin_bottom">25</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkImage">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="pixel_size">64</property>
                            <property name="icon_name">discord</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Community</property>
                            <attributes>
                              <attribute name="weight" value="semibold"/>
                              <attribute name="scale" value="1.1000000000000001"/>
                            </attributes>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">We meet, play, announce and
enjoy here. Feel free to join, 
ask question here</property>
                            <property name="justify">center</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">2</property>
                          </packing>
                        </child>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">True</property>
                    <property name="fill">True</property>
                    <property name="position">1</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkButton">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <signal name="clicked" handler="onChannelClicked" swapped="no"/>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="valign">center</property>
                        <property name="margin_left">10</property>
                        <property name="margin_right">10</property>
                        <property name="margin_start">10</property>
                        <property name="margin_end">10</property>
                        <property name="margin_top">25</property>
                        <property name="margin_bottom">25</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkImage">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="pixel_size">64</property>
                            <property name="icon_name">youtube</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">Channel</property>
                            <attributes>
                              <attribute name="weight" value="semibold"/>
                              <attribute name="scale" value="1.1000000000000001"/>
                            </attributes>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkLabel">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="label" translatable="yes">We upload guides, show-off
hacks and fixes
here</property>
                            <property name="justify">center</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">2</property>
                          </packing>
                        </child>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">True</property>
                    <property name="fill">True</property>
                    <property name="position">2</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">3</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="name">page4</property>
            <property name="title" translatable="yes">page4</property>
            <property name="position">4</property>
          </packing>
        </child>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="orientation">vertical</property>
            <child>
              <object class="GtkImage">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">12</property>
                <property name="margin_bottom">5</property>
                <property name="pixel_size">64</property>
                <property name="icon_name">dialog-information-symbolic</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">2</property>
                <property name="label" translatable="yes">Help improve rlxos</property>
                <attributes>
                  <attribute name="weight" value="heavy"/>
                  <attribute name="scale" value="1.5"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">2</property>
                <property name="label" translatable="yes">You can help rlxos for better support for hardware</property>
                <attributes>
                  <attribute name="weight" value="normal"/>
                  <attribute name="scale" value="1.2"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkBox">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="halign">center</property>
                <property name="valign">center</property>
                <property name="margin_bottom">150</property>
                <property name="orientation">vertical</property>
                <child>
                  <object class="GtkLabel">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="label" translatable="yes">Optional</property>
                    <attributes>
                      <attribute name="weight" value="ultrabold"/>
                      <attribute name="scale" value="1.2"/>
                    </attributes>
                  </object>
                  <packing>
                    <property name="expand">False</property>
                    <property name="fill">True</property>
                    <property name="position">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkLabel">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="margin_top">7</property>
                    <property name="margin_bottom">13</property>
                    <property name="label" translatable="yes">We have a very small team thats why its not possible for us
to test rlxos on every hardware. You can submit your system 
specifications that helps us to  imporve hardware support for rlxos</property>
                    <property name="justify">center</property>
                    <property name="wrap">True</property>
                  </object>
                  <packing>
                    <property name="expand">False</property>
                    <property name="fill">True</property>
                    <property name="position">1</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkButtonBox">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="layout_style">spread</property>
                    <child>
                      <object class="GtkButton">
                        <property name="label" translatable="yes">View Data</property>
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                        <property name="receives_default">True</property>
                        <signal name="clicked" handler="onViewDataClicked" swapped="no"/>
                      </object>
                      <packing>
                        <property name="expand">True</property>
                        <property name="fill">True</property>
                        <property name="position">0</property>
                      </packing>
                    </child>
                    <child>
                      <object class="GtkButton">
                        <property name="label" translatable="yes">Submit</property>
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                        <property name="receives_default">True</property>
                        <signal name="clicked" handler="onSubmitClicked" swapped="no"/>
                        <style>
                          <class name="suggested-action"/>
                        </style>
                      </object>
                      <packing>
                        <property name="expand">False</property>
                        <property name="fill">True</property>
                        <property name="position">1</property>
                      </packing>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">False</property>
                    <property name="fill">True</property>
                    <property name="position">2</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">3</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="name">page5</property>
            <property name="title" translatable="yes">page5</property>
            <property name="position">5</property>
          </packing>
        </child>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="halign">center</property>
            <property name="valign">center</property>
            <property name="margin_bottom">50</property>
            <property name="orientation">vertical</property>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">2</property>
                <property name="label" translatable="yes">Ready!</property>
                <attributes>
                  <attribute name="weight" value="heavy"/>
                  <attribute name="scale" value="1.5"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_top">8</property>
                <property name="margin_bottom">9</property>
                <property name="label" translatable="yes">rlxos is now ready for use</property>
                <attributes>
                  <attribute name="weight" value="normal"/>
                  <attribute name="scale" value="1.2"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton">
                <property name="label" translatable="yes">Enjoy</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <property name="halign">center</property>
                <property name="valign">center</property>
                <signal name="clicked" handler="onEnjoyClicked" swapped="no"/>
                <style>
                  <class name="suggested-action"/>
                </style>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="name">page6</property>
            <property name="title" translatable="yes">page6</property>
            <property name="position">6</property>
          </packing>
        </child>
      </object>
    </child>
    <child type="titlebar">
      <object class="GtkHeaderBar">
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <child>
          <object class="GtkButton" id="backButton">
            <property name="label" translatable="yes">Back</property>
            <property name="visible">True</property>
            <property name="can_focus">True</property>
            <property name="receives_default">True</property>
            <signal name="clicked" handler="onBackButtonClicked" swapped="no"/>
          </object>
        </child>
        <child>
          <object class="GtkButton" id="nextButton">
            <property name="label" translatable="yes">Next</property>
            <property name="visible">True</property>
            <property name="can_focus">True</property>
            <property name="receives_default">True</property>
            <signal name="clicked" handler="onNextButtonClicked" swapped="no"/>
            <style>
              <class name="suggested-action"/>
            </style>
          </object>
          <packing>
            <property name="pack_type">end</property>
            <property name="position">1</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
</interface>`
